<?php

// include_once "";

class BeanstalkTask extends \Phalcon\Cli\Task
{
    public $bst;

    public function initialize(){
        $this->bst = new \Beanstalkd($host='192.168.47.125', $port='11300', $tube='default');
    }

    public function mainAction(){
        $this->helpAction();
    }

    public function helpAction(){
        echo "beanstalk [cmd]".PHP_EOL;
        echo "[cmd] ... ".PHP_EOL;
        echo "  stats                   : see all q in beanstalk server".PHP_EOL;
        echo "  version                 : ".PHP_EOL;
        echo "  help                    : ".PHP_EOL;
        echo "  putjob [params|...]     : ".PHP_EOL;
        echo "  getjob                  : ".PHP_EOL;
        echo "  deljob [id]             : deljob 1 ".PHP_EOL;
    }

    public function versionAction(){
        var_dump($this->bst->version());
    }

    public function statsAction(){
        var_dump($this->bst->stats());
    }

    public function putjobAction($params){
        // $jobContent = [
        //     "name" => "max",
        //     "position" => "developer"
        // ];
        $jobId = $this->bst->putJob($params);
        var_dump($jobId);
    }

    public function getjobAction(){
        $job = $this->bst->getJobByReady();
        var_dump($job);
    }

    public function deljobAction($params){
        $jobId = $params[0];

        if (empty($jobId)) {
            echo "job id empty";
            return;
        }

        $this->bst->delJobByJobId($jobId);
    }

}
