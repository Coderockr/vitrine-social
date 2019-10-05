import React from 'react';
import { Upload, Icon, Modal } from 'antd';
import BottomNotification from '../BottomNotification';
import styles from './styles.module.scss';

class UploadImages extends React.Component {
  constructor(props) {
    super(props);

    const fileList = props.images ? props.images.map(image => ({
      uid: image.id,
      name: image.name,
      status: 'done',
      url: image.url,
    })) : [];

    this.state = {
      fileList,
      previewVisible: false,
      previewImage: '',
      updateList: [],
    };
  }

  getBase64(img, callback) {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result));
    reader.readAsDataURL(img);
  }

  beforeUpload(file) {
    const isValidSize = file.size / 1024 / 1024 < 2;
    if (!isValidSize) {
      BottomNotification({ message: 'Imagem é muito grande. Deve ter menos de 2MB.', success: false });
      return false;
    }
    const isValidFormat = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isValidFormat) {
      BottomNotification({ message: 'Imagem com formato inválido. Deve ser JPEG ou PNG.', success: false });
      return false;
    }
    return true;
  }

  handleCancel = () => this.setState({ previewVisible: false })

  handlePreview = (file) => {
    this.setState({
      previewImage: file.url || file.thumbUrl,
      previewVisible: true,
    });
  }

  handleChange = (info) => {
    this.setState({
      fileList: info.fileList,
    });
  }

  removeImage(item) {
    const find = this.state.updateList.find(image => image.file.uid === item.uid);
    let updateList = [...this.state.updateList];
    if (find) {
      const index = updateList.indexOf(find);
      updateList.splice(index, 1);
    } else {
      const updateFile = { file: item, action: 'delete' };
      updateList = [...updateList, updateFile];
    }
    this.setState({ updateList });
    this.props.onChange(updateList);
  }

  uploadImage({ onSuccess, file }) {
    this.getBase64(file, (imageUrl) => {
      const updateFile = { file, action: 'add', imageUrl };
      const updateList = [...this.state.updateList, updateFile];
      this.setState({ updateList });
      this.props.onChange(updateList);
      onSuccess(null, updateFile);
    });
  }

  render() {
    const { previewVisible, previewImage, fileList } = this.state;
    const uploadButton = (
      <div>
        <Icon type="plus" />
        <div className="ant-upload-text">Adicionar</div>
      </div>
    );
    return (
      <div className={styles.uploadWrapper}>
        <Upload
          listType="picture-card"
          fileList={fileList}
          customRequest={({ onSuccess, onError, file }) => this.uploadImage({
            onSuccess,
            onError,
            file,
          })}
          beforeUpload={this.beforeUpload}
          onPreview={this.handlePreview}
          onChange={this.handleChange}
          onRemove={item => this.removeImage(item)}
        >
          {uploadButton}
        </Upload>
        <Modal visible={previewVisible} footer={null} onCancel={this.handleCancel}>
          <img
            className={styles.modalImage}
            alt="example"
            style={{ width: '100%' }}
            src={previewImage}
          />
        </Modal>
      </div>
    );
  }
}

export default UploadImages;
