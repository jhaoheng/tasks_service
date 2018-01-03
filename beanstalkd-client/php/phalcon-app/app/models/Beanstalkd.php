<?php  

use Phalcon\Queue\Beanstalk;
/**
* 
*/
class Beanstalkd
{
    public $_queue;
    
    function __construct($host='localhost', $port='11300', $tube='default')
    {
        try {
            // Connect to the queue
            $queue = new Beanstalk(
                [
                    "host" => $host,
                    "port" => $port,
                    'persistent' => true,
                ]
            );
            // 
            $queue->choose($tube);
            $queue->connect();
            $this->_queue = $queue;
        } catch (Exception $e) {
            die($e->getMessage().PHP_EOL.PHP_EOL);
        }
    }

    public function close(){
        $queue = $this->_queue;
        $queue->quit();
    }

    public function version(){
        $queue = $this->_queue;
        $stats = $queue->stats();
        return $stats['version'];
    }

    public function stats(){
        $queue = $this->_queue;
        $stats = $queue->stats();
        return $stats;
    }


    public function putJob($jobContent, $delay=1){
        $queue = $this->_queue;

        $attributes = array('priority'  => 0,
                            'delay'     => $delay, 
                            // "ttr"       => 3600
                            );

        $jobId = $queue->put($jobContent, $attributes);
        return $jobId;
    }

    public function getJobByReady(){
        $queue = $this->_queue;
        $job = $queue->reserve(); // 獲取並鎖住任務

        $jobInfo = [
            "id" => $job->getId(),
            "body" => $job->getBody()
        ];
        return $jobInfo;
    }

    public function getJobByJobId($jobId){
        $queue = $this->_queue;
        $job = $queue->jobPeek($jobId);
        $this->_job = $job;
        return $job;
    }

    public function delJobByJobId($jobId){
        $job = $this->getJobByJobId($jobId);
        $job->delete();
    }

    // 取得 tube 上的資料狀態
    public function getTubeStat(){
        $queue = $this->_queue;
        $data = $queue->statsTube($tube);
        return $data;

        /* key
        name
        current-jobs-urgent
        current-jobs-ready
        current-jobs-reserved
        current-jobs-delayed
        current-jobs-buried
        total-jobs
        current-using
        current-watching
        current-waiting
        cmd-delete
        cmd-pause-tube
        pause
        pause-time-left
        */
    }
}


?>