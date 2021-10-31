// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "./log.sol";

contract SimplePaymentChannel is ChainlinkClient, Console {
    //chainlink init
    using Chainlink for Chainlink.Request;

    event taskCommit(string dockerImage, string port,uint transactionId);

    bytes32 public finished;
    uint public count;
    string public result;
    enum State {Waiting, Failed, Succeeded, Canceled}

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
        // 4->Canceled
        State state;
        JobInformation jobs;
        uint creationTimeStamp;
    }

    //customer to transactions
    mapping(address => Transaction[]) public customerToTransactions;
    // transactionId to transaction
    mapping(uint => Transaction) public tidToTransaction;
    //requestId to Transaction
    mapping(bytes32 => uint) public requestIdToTransactionId;

    // array below will be used to iterate above mapping
    address[] public consumers;


    constructor ()
    {
        // set Chainlink configuration(to enable a HTTP-GET request later)
        setPublicChainlinkToken();
    }

    // customer calls this function to publish task
    function publishTask(
        string calldata dockerImage,
        string calldata port,
        bytes32 flagMessage,
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
        if (customerToTransactions[msg.sender].length == 0) {
            consumers.push(msg.sender);
        }
        customerToTransactions[msg.sender].push(transaction);
        emit taskCommit(dockerImage, port, transaction.transactionId);
    }

    // retrieve all tasksInfo from a target user
    function retrieveTasksInfoFromTargetUser(address consumerAddress) external view returns (
        uint[] memory transactionIds,
    // we need to reflect this value manually in our client
        State[] memory states,
        string[] memory dockerImages,
        string[] memory ports,
        bytes32[] memory flagMessages,
        string[] memory urls,
        uint[] memory creationTimeStamps
    ){

        Transaction[] memory transactions = customerToTransactions[consumerAddress];

        transactionIds = new uint[](transactions.length);
        states = new State[](transactions.length);
        dockerImages = new string[](transactions.length);
        ports = new string[](transactions.length);
        flagMessages = new bytes32[](transactions.length);
        urls = new string[](transactions.length);
        creationTimeStamps = new uint[](transactions.length);

        for (uint i = 0; i < transactions.length; i++) {
            transactionIds[i] = transactions[i].transactionId;
            states[i] = transactions[i].state;
            dockerImages[i] = transactions[i].jobs.dockerImage;
            ports[i] = transactions[i].jobs.port;
            flagMessages[i] = transactions[i].jobs.flagMessage;
            urls[i] = transactions[i].jobs.url;
            creationTimeStamps[i] = transactions[i].creationTimeStamp;
        }


        return (transactionIds, states, dockerImages, ports, flagMessages, urls, creationTimeStamps);
    }

    // retrieve all unfinished(!= succeeded) state tasks
    // cause the returns should not return structure, so I use 2 arrays to return
    function retrieveAllUnfinishedTask() external view returns (
        string[] memory dockerImages,
        string[] memory ports,
        uint[] memory transactionIds){
        uint total = 0;
        for (uint index = 0; index < consumers.length; index++) {
            for (uint j = 0; j < customerToTransactions[consumers[index]].length; j++) {
                if (customerToTransactions[consumers[index]][j].state == State.Failed || customerToTransactions[consumers[index]][j].state == State.Waiting) {
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
                if (customerToTransactions[consumers[index]][j].state == State.Failed || customerToTransactions[consumers[index]][j].state == State.Waiting) {
                    dockerImages[flag] = customerToTransactions[consumers[index]][j].jobs.dockerImage;
                    ports[flag] = customerToTransactions[consumers[index]][j].jobs.port;
                    transactionIds[flag] = customerToTransactions[consumers[index]][j].transactionId;
                    flag++;
                }
            }
        }

        return (dockerImages, ports, transactionIds);
    }


    // customer calls this function and return the money to customer
    function returnMoneyBack(uint transactionId) external payable {
        // every task will cost about 0.01eth
        // find this transaction
        Transaction storage transaction = tidToTransaction[transactionId];
        // check if the sender and the customer is the same person
        // check the time, the customer only allowed to withdraw his/her deposit when the time is out(>1H)
        //TODO change the block.time into real-world time using chainlink
        require(block.timestamp > transaction.creationTimeStamp + 1*60*60);
        // check the status of this task
        require(transaction.state == State.Waiting || transaction.state == State.Failed);

        transaction.state = State.Canceled;
        payable(transaction.customer).transfer(transaction.money);
    }

    // cloud provider commit his tasks(only one task is allowed)
    function commitTask(string calldata ipAddress, uint transactionId) external {
        // check the transactionId is available and set the requestId into it
        require(tidToTransaction[transactionId].creationTimeStamp != 0);
        tidToTransaction[transactionId].cloudProvider = payable(msg.sender);


        // using Liveness Probe to check the running status of image
        string memory url =string(abi.encodePacked("http://", ipAddress, ":", tidToTransaction[transactionId].jobs.port, tidToTransaction[transactionId].jobs.url));

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
        finished = flagMessage;
        uint tid = requestIdToTransactionId[_requestId];
        Transaction storage transaction = tidToTransaction[tid];
        if (transaction.jobs.flagMessage == flagMessage) {
            transaction.state = State.Succeeded;
            //send money to the cloud provider
            payable(transaction.cloudProvider).transfer(transaction.money);
        } else {
            transaction.state = State.Failed;
        }
    }
}