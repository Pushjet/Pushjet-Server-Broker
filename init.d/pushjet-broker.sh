#! /bin/sh
# /etc/init.d/pushjet-broker


PIDFILE=/var/run/pushjet-broker.pid
NAME=pushjet-broker
BIN=/opt/pushjet/pushjet-broker

case $1 in
  start)
    echo "Starting $NAME"
    start-stop-daemon --start --exec $BIN -m --pidfile $PIDFILE --background -d /etc/
    ;;
  stop)
    echo "Stopping $NAME"
    start-stop-daemon --stop --pidfile $PIDFILE
    ;;
  restart)
    echo "Restarting $NAME"
    $0 stop
    sleep 1
    $0 start
    ;;
  status)
    start-stop-daemon --status --pidfile $PIDFILE
    status=$?
    if [ $status -eq 0 ]; then
       echo "$NAME is running"
    elif [ $status -eq 1 ]; then
       echo "$NAME is not running but the PID file exists"
    elif [ $status -eq 3 ]; then
       echo "$NAME is not running"
    elif [ $status -eq 4 ]; then
       echo "Can't determine if $NAME is running"
    fi
    ;;
  *)
    echo "Usage: $0 {start|stop|restart|status}"
    exit 1
    ;;
esac

exit 0

