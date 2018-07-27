import { notification } from 'antd';
import cx from 'classnames';
import styles from './styles.module.scss';

const BottomNotification = ({ message, success, onClose }) => {
  notification.config({
    placement: 'bottomRight',
    bottom: 0,
  });
  notification.open({
    message,
    className: cx(styles.notification, { [styles.success]: success }),
    duration: onClose ? 0 : 3,
    onClose,
  });
};

export default BottomNotification;
