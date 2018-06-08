import { notification } from 'antd';
import styles from './styles.module.scss';

const BottomNotification = (message) => {
  notification.config({
    placement: 'bottomRight',
    bottom: 0,
  });
  notification.open({
    message,
    className: styles.notification,
    duration: 3,
  });
};

export default BottomNotification;
