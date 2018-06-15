import { notification } from 'antd';
import cx from 'classnames';
import styles from './styles.module.scss';

const BottomNotification = ({ message, success }) => {
  notification.config({
    placement: 'bottomRight',
    bottom: 0,
  });
  notification.open({
    message,
    className: cx(styles.notification, { [styles.success]: success }),
    duration: 3,
  });
};

export default BottomNotification;
