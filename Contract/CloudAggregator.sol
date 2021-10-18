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
        string flagMessage;
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

    mapping(address => Transaction[]) customerToTransactions;
    mapping(uint => Transaction) tidToTransaction;
    // array below will be used to iterate above mapping
    address[] consumers;


    constructor (address payable recipientAddress)
    payable
    {
        // set Chainlink configuration(to enable a HTTP-GET request later)
        setPublicChainlinkToken();
        oracle = 0xc57B33452b4F7BB189bB5AfaE9cc4aBa1f7a4FD8;
        jobId = "d5270d1c311941d0b08bead21fea7747";
        fee = 0.1 * 10 ** 18;


        Transaction transaction;
        transaction.customer = msg.sender;

        sender = payable(msg.sender);
        recipient = recipientAddress;
        // block.timestamp is the current block in seconds since the epoch
        // only after 24H and the task is not finished, the customer could take back his/her deposit
        expiration = block.timestamp + 24 * 60 * 60;
    }

    // customer calls this function to publish task
    function publishTask(
        string dockerImage,
        string port,
        string flagMessage,
        string url,
    //TODO change this param into decentralized type
        uint creationTimestamp
    ) external payable {
        Transaction transaction;
        transaction.customer = msg.sender;
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
    function retrieveTasksInfoFromTargetUser(address consumerAddress, uint index) external returns (
        string transactionId,
    // we need to reflect this value manually in our client
        State state,
        string dockerImage,
        string port,
        string flagMessage,
        string url,
        uint creationTimeStamp
    ){
        //check the length
        require(index < customerToTransactions[consumerAddress].length);

        return (customerToTransactions[consumerAddress][index].transactionId,
        customerToTransactions[consumerAddress][index].state,
        customerToTransactions[consumerAddress][index].jobs.dockerImage,
        customerToTransactions[consumerAddress][index].jobs.port,
        customerToTransactions[consumerAddress][index].jobs.flagMessage,
        customerToTransactions[consumerAddress][index].jobs.url,
        customerToTransactions[consumerAddress][index].creationTimeStamp);
    }

    //related to above function
    function getUserTasksLength(address consumerAddress) external returns (uint index){
        return customerToTransactions[consumerAddress].length;
    }

    // retrieve one the unfinished(!= succeeded) state task
    // cause the returns should not return structure, so I use 2 arrays to return
    function retrieveOneUnfinishedTask() external returns (string[] dockerImage, string[] transactionId){
        for (uint index = 0; index < consumers.length; index++) {
            for (uint j = 0; j < customerToTransactions[consumers[index]].length; j++) {
                if (customerToTransactions[consumers[index]][j].state != State.Succeeded) {
                    dockerImage.push(customerToTransactions[consumers[index]][j].jobs.dockerImage);
                    transactionId.push(customerToTransactions[consumers[index]][j].transactionId);
                }
            }
        }
        return;
    }


    // customer calls this function and return the money to customer
    function returnMoneyBack(uint transactionId) external payable {
        // find this transaction
        Transaction[] transactions = customerToTransactions[msg.sender];
        uint transactionsLength = transactions.length;
        uint index = 0;
        while (index < transactionsLength) {
            if (transactions[i].transactionId == transactionId) {
                break;
            }
            index++;
        }
        // check the time, the customer only allowed to withdraw his/her deposit when the time is out(>24H)
        //TODO change the block.time into real-world time using chainlink
        require(block.time > transactions[index].creationTimeStamp + 24 * 60 * 60);
        // check the status of this task
        require(transactions[index].state != State.Succeeded);

        payable(msg.sender).transfer();
    }

    // cloud provider commit his tasks(only one task is allowed)
    function commitTask(string ipAddress, uint transactionId) external {
        // using Liveness Probe to check the running status of image
        string memory url = new string("http://" + ipAddress + ":" + tidToTransaction[transactionId].jobs.port + tidToTransaction[transactionId].jobs.url);

        //TODO check if it's available to get the message
        Chainlink.Request memory request = buildChainlinkRequest(jobId, address(this), this.fulfillLivenessCheck.selector);


        // Set the URL to perform the GET request on
        request.add("get", url);

        // Set the path to find the desired data in the API response, where the response format is:
        // {
        //    "keyword":"<keyword of Liveness probe>"
        //  }
        request.add("path", "keyword");

        // Multiply the result by 1000000000000000000 to remove decimals
        int timesAmount = 10 ** 18;
        request.addInt("times", timesAmount);

        // Sends the request
        return sendChainlinkRequestTo(oracle, request, fee);

    }

    function fulfillLivenessCheck(bytes32 _requestId, uint256 _volume) public recordChainlinkFulfillment(_requestId)
    {
        volume = _volume;
    }
}