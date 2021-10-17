// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract SimplePaymentChannel {
    uint count;
    enum State {Waiting, Failed, Succeeded}

    // the Information of Job
    struct JobInformation {
        // dockerImage label
        string dockerImage;
        // Liveness Probe(port + message)
        uint port;
        string flagMessage;
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
    // array below will be used to iterate above mapping
    address[] consumers;


    constructor (address payable recipientAddress)
    payable
    {
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
        uint port,
        string flagMessage,
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

        //creation timestamp
        transaction.creationTimeStamp = creationTimestamp;
        //TODO generate a distributed taskId
        transaction.transactionId = creationTimestamp + count;
        count = count + 1;

        //add consumer to the address[] consumers
        if (customerToTransactions[msg.sender].length != 0) {
            consumers.push(msg.sender);
        }
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

    // retrieve one the unfinished(!= succeeded) state task
    // cause the returns should not return structure, so I use 2 arrays to return
    function retrieveOneUnfinishedTask() external returns (string[] dockerImage, string[] transactionId){
        for (uint index = 0; index < consumers.length; index++) {
            for (uint j = 0; j < customerToTransactions[consumers[index]].length; j++) {
                if (customerToTransactions[consumers[index]][j].state != State.Succeeded){
                    dockerImage.push(customerToTransactions[consumers[index]][j].jobs.dockerImage);
                    transactionId.push(customerToTransactions[consumers[index]][j].transactionId);
                }
            }
        }
        return;
    }

    //
}