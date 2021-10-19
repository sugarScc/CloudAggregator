// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";

contract SimplePaymentChannel is ChainlinkClient {
    //chainlink init
    using Chainlink for Chainlink.Request;


    uint count;
    enum State {Waiting, Failed, Succeeded}

    // the Information of Job
    struct JobInformation {
        // dockerImage label
        string dockerImage;
        // Liveness Probe(port + message), see LivenessProbe in Kubernetes
        string port;
        // the chainlink will always return bytes32, so we just store the bytes32 one
        bytes32 flagMessage;
        string url;

        //IP address of the docker instance
        string ipAddress;
    }


    struct Transaction {
        //TODO generate this transactionId by VRF(chainlink)
        uint transactionId;
        // the address of customer
        address payable customer;
        // the address of cloudProvider
        address payable cloudProvider;
        // the amount of money of this transaction
        uint256 money;
        // status
        // 1->failed
        // 2->waiting
        // 3->succeeded
        State state;
        JobInformation jobs;
        uint creationTimeStamp;
    }

    //customer to transactions
    mapping(address => Transaction[]) customerToTransactions;
    // transactionId to transaction
    mapping(uint => Transaction) tidToTransaction;
    //requestId to Transaction
    mapping(bytes32 => uint) requestIdToTransactionId;

    // array below will be used to iterate above mapping
    address[] consumers;


    constructor ()
    {
        // set Chainlink configuration(to enable a HTTP-GET request later)
        setPublicChainlinkToken();
    }

    // customer calls this function to publish task
    function publishTask(
        string calldata dockerImage,
        string calldata port,
        bytes32  flagMessage,
        string calldata url,
    //TODO change this param(creationTimestamp) into decentralized type
        uint creationTimestamp
    ) external payable {
        Transaction memory transaction;
        transaction.customer = payable(msg.sender);
        transaction.money = msg.value;
        transaction.state = State.Waiting;
        transaction.jobs.dockerImage = dockerImage;
        transaction.jobs.flagMessage = flagMessage;
        transaction.jobs.port = port;
        transaction.jobs.url = url;

        //creation timestamp
        transaction.creationTimeStamp = creationTimestamp;
        //TODO generate a distributed taskId
        transaction.transactionId = creationTimestamp + count;
        count = count + 1;
        // init reverse index
        tidToTransaction[transaction.transactionId] = transaction;


        //add consumer to the address[] consumers
        if (customerToTransactions[msg.sender].length != 0) {
            consumers.push(msg.sender);
        }
    }

    // retrieve tasksInfo from a target user
    // to reduce gas, this function need to be called by client as iterator
    function retrieveTasksInfoFromTargetUser(address consumerAddress, uint index) external view returns (
        uint transactionId,
    // we need to reflect this value manually in our client
        State state,
        string memory dockerImage,
        string memory port,
        bytes32 flagMessage,
        string memory url,
        uint creationTimeStamp
    ){
        //check the length
        require(index < customerToTransactions[consumerAddress].length);

        Transaction memory transaction = customerToTransactions[consumerAddress][index];

        return (transaction.transactionId,
        transaction.state,
        transaction.jobs.dockerImage,
        transaction.jobs.port,
        transaction.jobs.flagMessage,
        transaction.jobs.url,
        transaction.creationTimeStamp);
    }

    //related to above function
    function getUserTasksLength(address consumerAddress) external view returns (uint index){
        return customerToTransactions[consumerAddress].length;
    }

    // retrieve one the unfinished(!= succeeded) state task
    // cause the returns should not return structure, so I use 2 arrays to return
    function retrieveOneUnfinishedTask() external view returns (
        string[] memory dockerImages,
        string[] memory ports,
        uint[] memory transactionIds){
        uint total = 0;
        for (uint index = 0; index < consumers.length; index++) {
            for (uint j = 0; j < customerToTransactions[consumers[index]].length; j++) {
                if (customerToTransactions[consumers[index]][j].state != State.Succeeded) {
                    total++;
                }
            }
        }

        dockerImages = new string[](total);
        ports = new string[](total);
        transactionIds = new uint[](total);
        uint flag = 0;
        for (uint index = 0; index < consumers.length; index++) {
            for (uint j = 0; j < customerToTransactions[consumers[index]].length; j++) {
                if (customerToTransactions[consumers[index]][j].state != State.Succeeded) {
                    dockerImages[flag] = customerToTransactions[consumers[index]][j].jobs.dockerImage;
                    ports[flag] = customerToTransactions[consumers[index]][j].jobs.port;
                    transactionIds[flag] = customerToTransactions[consumers[index]][j].transactionId;
                }
            }
        }

        return (dockerImages, ports, transactionIds);
    }


    // customer calls this function and return the money to customer
    function returnMoneyBack(uint transactionId) external payable {
        // find this transaction
        Transaction[] memory transactions = customerToTransactions[msg.sender];
        Transaction memory transaction;
        uint transactionsLength = transactions.length;
        uint index = 0;
        while (index < transactionsLength) {
            if (transactions[index].transactionId == transactionId) {
                transaction = transactions[index];
            }
            index++;
        }
        // check the time, the customer only allowed to withdraw his/her deposit when the time is out(>24H)
        //TODO change the block.time into real-world time using chainlink
        require(block.timestamp > transaction.creationTimeStamp + 24 * 60 * 60);
        // check the status of this task
        require(transaction.state != State.Succeeded);

        payable(msg.sender).transfer(transaction.money);
    }

    // cloud provider commit his tasks(only one task is allowed)
    function commitTask(string calldata ipAddress, uint transactionId) external {
        // check the transactionId is available and set the requestId into it
        require(tidToTransaction[transactionId].creationTimeStamp != 0);
        tidToTransaction[transactionId].cloudProvider = payable(msg.sender);


        // using Liveness Probe to check the running status of image
        string memory url =  string(abi.encodePacked("http://", ipAddress,":", tidToTransaction[transactionId].jobs.port, tidToTransaction[transactionId].jobs.url));

        //TODO check if it's available to get the message
        Chainlink.Request memory request = buildChainlinkRequest("7401f318127148a894c00c292e486ffd", address(this), this.fulfillLivenessCheck.selector);


        // Set the URL to perform the GET request on
        request.add("get", url);
        request.add("path", "keyword");

        // Sends the request
        bytes32 requestId = sendChainlinkRequestTo(0xc57B33452b4F7BB189bB5AfaE9cc4aBa1f7a4FD8, request, 3 * 10 ** 18);
        requestIdToTransactionId[requestId] = transactionId;
        return;
    }

    // receive the response from above commitTask, need to check if it's correct response and change the state of this task
    function fulfillLivenessCheck(bytes32 _requestId, bytes32 flagMessage) public recordChainlinkFulfillment(_requestId)
    {
        uint tid = requestIdToTransactionId[_requestId];
        Transaction memory transaction = tidToTransaction[tid];
        if (transaction.jobs.flagMessage == flagMessage) {
            transaction.state = State.Succeeded;
            //send money to the cloud provider
            payable(transaction.cloudProvider).transfer(transaction.money);
        } else{
            transaction.state = State.Failed;
        }
    }
}