#!/usr/bin/env bash
exe_path="$( cd "$( dirname "$0"  )" && pwd  )"
WebSocket_HOME=${exe_path}/../
WebSocket_SBIN=${WebSocket_HOME}/bin/
WebSocket_PID=${WebSocket_HOME}/data/websocket.pid

start()
{
    if [ -f ${WebSocket_PID} ]
    then
        pid=`cat ${WebSocket_PID}`
        process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep "websocket" | wc -l`
        if [ ${process_num} -ge 1 ];
        then
            echo "service already running. pid=$pid"
            return
        fi
    fi
    cd ${WebSocket_SBIN}
    nohup ./websocket &> ../logs/websocket.log 2>> ../logs/websocket_except.log &
    echo "websocket start"
}


stop()
{
    if [ ! -f ${WebSocket_PID} ]
    then
        echo "service already exit"
        return
    fi
    pid=`cat ${WebSocket_PID}`
    process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep "websocket" | wc -l`
    if [ ${process_num} -eq 0 ];
    then
        echo "service already exit  $process_num"
        return
    fi
    kill -TERM `cat ${WebSocket_PID}`
    ret=$?
    if [ ${ret} -eq 0 ]
    then
        echo "websocket stop"
    else
        echo "websocket stop failed"
    fi
    return
}

restart()
{
    stop
    start
    return
}


reload()
{
    if [ ! -f ${WebSocket_PID} ]
    then
        echo "service already exit"
        return
    fi
    pid=`cat ${WebSocket_PID}`
    process_num=`ps -ef | grep -w ${pid} | grep -v "grep" | grep "websocket" | wc -l`
    if [ ${process_num} -eq 0 ];
    then
        echo "service already exit"
        return
    fi
    kill -USR2 `cat $${WebSocket_PID}`
    return
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    reload)
        reload
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart|reload}"
        exit 1
esac

exit 0