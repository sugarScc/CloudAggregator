// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "./log.sol";

contract SimplePaymentChannel is ChainlinkClient, Console {
    //chainlink init
    using Chainlink for Chainlink.Request;

    event taskCommit(string dockerImage, string port, uint transactionId);

    bytes32 public finished;
    uint public count;
    string public result;
    address[] public providers;
    mapping(address => bytes32) public providersToName;
    mapping(bytes32 => uint) public providersToSLA;

    mapping(address => uint) public rewardedUser;
    //the Metrics is network latency. and the SLA is only from 1-5
    mapping(bytes32 => uint[]) public providerSLAMetrics;
    MetricsInformation[] public dockerToProviderMetrics;
    mapping(bytes32 => address) public dockerToProvider;
    mapping(bytes32 => uint) public providerCalRound;
    mapping(bytes32 => uint) public providerDeposit;
    uint public totalRewardPool;

    struct MetricsInformation {
        // dockerImage label
        address metricsProvider;
        bytes32 AimCloudProvider;
        uint metrics;
        // which round do this metrics provider engaged in
        uint round;
    }



    enum State {Waiting, Failed, Succeeded, Canceled}

    // the Information of Job
    struct JobInformation {
        // dockerImage label
        string dockerImage;
        // Liveness Probe(port + message), see LivenessProbe in Kubernetes
        string port;
        // the chainlink will always return bytes32, so we just store the bytes32 one
        bytes32 flagMessage;
        string path;

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
        transaction.jobs.path = url;

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
        string[] memory path,
        string[] memory urls,
        uint[] memory creationTimeStamps
    ){

        Transaction[] memory transactions = customerToTransactions[consumerAddress];

        transactionIds = new uint[](transactions.length);
        states = new State[](transactions.length);
        dockerImages = new string[](transactions.length);
        ports = new string[](transactions.length);
        flagMessages = new bytes32[](transactions.length);
        path = new string[](transactions.length);
        creationTimeStamps = new uint[](transactions.length);
        urls = new string[](transactions.length);

        for (uint i = 0; i < transactions.length; i++) {
            transactionIds[i] = transactions[i].transactionId;
            states[i] = transactions[i].state;
            urls[i] = transactions[i].jobs.url;
            dockerImages[i] = transactions[i].jobs.dockerImage;
            ports[i] = transactions[i].jobs.port;
            flagMessages[i] = transactions[i].jobs.flagMessage;
            path[i] = transactions[i].jobs.path;
            creationTimeStamps[i] = transactions[i].creationTimeStamp;
        }


        return (transactionIds, states, dockerImages, ports, flagMessages, path, urls, creationTimeStamps);
    }

    // retrieve all unfinished(!= succeeded) state taskss
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

    function providerRegister(bytes32 providerName, uint SLALevel) external {
        //用于注册云服务商
        providersToName[msg.sender] = providerName;
        providersToSLA[providerName] = SLALevel;
    }

    function metricsCollection(string calldata dockerImage, bytes32 providerName, uint SLAMetrics) public payable{
        //for now, the threshold is 100
        require(providerSLAMetrics[providerName].length <= 100);
        providerSLAMetrics[providerName].push(SLAMetrics);
        uint total = 0;
        if (providerSLAMetrics[providerName].length == 100) {
            //calculate the SLAMetrics and adjust it
            for (uint index = 0; index < 100; index++) {
                total = total + providerSLAMetrics[providerName][index];
            }
            uint averageLatency = total / 100;
            if (averageLatency > 1000) {
                total = 5;
            } else if (averageLatency > 500) {
                total = 4;
            } else if (averageLatency > 200) {
                total = 3;
            } else if (averageLatency > 100) {
                total = 2;
            } else if (averageLatency > 50) {
                total = 1;
            }
            delete providerSLAMetrics[providerName];
            //deduct the deposit for all lying provider
            if (providersToSLA[providerName] < total) {
                if (total == 2) {
                    totalRewardPool += providerDeposit[providerName] - 14;
                }
                if (total == 3) {
                    totalRewardPool += providerDeposit[providerName] - 5;
                }
                if (total == 4) {
                    totalRewardPool += providerDeposit[providerName] - 2;
                }
                if (total == 5) {
                    totalRewardPool += providerDeposit[providerName] - 1;
                }
                providersToSLA[providerName] = total;
            }
            uint rewardAmount = totalRewardPool / 10000;
            //calculate for all the eligible user to get the reward
            for (uint index = 0; index < dockerToProviderMetrics.length; index++) {
                if (dockerToProviderMetrics[index].round == providerCalRound[providerName]) {
                    if ((dockerToProviderMetrics[index].metrics*7)/ 10 > averageLatency  && (dockerToProviderMetrics[index].metrics*13) /10 < averageLatency) {
                        payable(dockerToProviderMetrics[index].metricsProvider).transfer(rewardAmount);
                    }
                }
            }
            providerCalRound[providerName] = providerCalRound[providerName] + 1;
        }
    }

    // customer calls this function and return the money to customer
    function returnMoneyBack(uint transactionId) external payable {
        // every task will cost about 0.01eth
        // find this transaction in tidToTransaction Map
        Transaction storage transaction = tidToTransaction[transactionId];
        // check if the sender and the customer is the same person
        // check the time, the customer only allowed to withdraw his/her deposit when the time is out(>1H)
        //TODO change the block.time into real-world time using chainlink
        require(block.timestamp > transaction.creationTimeStamp + 1 * 60 * 60);
        // check the status of this task
        require(transaction.state == State.Waiting || transaction.state == State.Failed);

        transaction.state = State.Canceled;
        // find this transaction in customerToTransactions
        Transaction[] memory transactions = customerToTransactions[transaction.customer];
        uint index = 0;
        for (; index < transactions.length; index++) {
            if (transactions[index].transactionId == transactionId) {
                break;
            }
        }
        Transaction storage transactionInMap2 = customerToTransactions[transaction.customer][index];
        transactionInMap2.state = State.Canceled;

        payable(transaction.customer).transfer(transaction.money);
    }

    // cloud provider commit his tasks(only one task is allowed)
    function commitTask(string calldata ipAddress, uint transactionId) external {
        // check the transactionId is available and set the requestId into it
        require(tidToTransaction[transactionId].creationTimeStamp != 0);
        // find the 1st ranked provider
        address _1stProvider = providers[0];
        uint index;
        for (index = 0; index < providers.length; index++) {
            if (providersToSLA[providersToName[providers[index]]] > providersToSLA[providersToName[_1stProvider]]) {
                _1stProvider = providers[index];
            }
        }
        //这个地方要求只能是排名第一的云服务商来参与部署镜像
        require(_1stProvider == msg.sender);
        // we prevent any unqualified cloud service provider in here
        tidToTransaction[transactionId].cloudProvider = payable(msg.sender);

        // find this transaction in customerToTransactions
        Transaction[] memory transactions = customerToTransactions[tidToTransaction[transactionId].customer];
        for (; index < transactions.length; index++) {
            if (transactions[index].transactionId == transactionId) {
                break;
            }
        }
        Transaction storage transactionInMap2 = customerToTransactions[tidToTransaction[transactionId].customer][index];
        transactionInMap2.cloudProvider = payable(msg.sender);


        // using Liveness Probe to check the running status of image
        string memory url = string(abi.encodePacked("http://", ipAddress, ":", tidToTransaction[transactionId].jobs.port, tidToTransaction[transactionId].jobs.path));
        transactionInMap2.jobs.url = url;
        tidToTransaction[transactionId].jobs.url = url;

        //TODO check if it's available to get the message
        Chainlink.Request memory request = buildChainlinkRequest("7401f318127148a894c00c292e486ffd", address(this), this.fulfillLivenessCheck.selector);


        // Set the URL to perform the GET request on
        request.add("get", url);
        request.add("path", "keyword");

        // Sends the request
        bytes32 requestId = sendChainlinkRequestTo(0xc57B33452b4F7BB189bB5AfaE9cc4aBa1f7a4FD8, request, 10 * 10 ** 18);
        requestIdToTransactionId[requestId] = transactionId;
        return;
    }

    // receive the response from above commitTask, need to check if it's correct response and change the state of this task
    function fulfillLivenessCheck(bytes32 _requestId, bytes32 flagMessage) public recordChainlinkFulfillment(_requestId)
    {
        finished = flagMessage;
        uint tid = requestIdToTransactionId[_requestId];
        Transaction storage transaction = tidToTransaction[tid];
        // find this transaction in customerToTransactions
        Transaction[] memory transactions = customerToTransactions[transaction.customer];
        uint index = 0;
        for (; index < transactions.length; index++) {
            if (transactions[index].transactionId == transaction.transactionId) {
                break;
            }
        }
        Transaction storage transactionInMap2 = customerToTransactions[transaction.customer][index];

        if (transaction.jobs.flagMessage == flagMessage) {
            transaction.state = State.Succeeded;
            transactionInMap2.state = State.Succeeded;
            //send money to the cloud provider
            payable(transaction.cloudProvider).transfer(transaction.money);
        } else {
            transaction.state = State.Failed;
            transaction.jobs.url = "";
            transactionInMap2.state = State.Failed;
            transactionInMap2.jobs.url = "";

        }
    }
}