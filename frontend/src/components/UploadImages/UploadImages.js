import React from 'react';
import { Upload, Icon, Modal } from 'antd';
import styles from './styles.module.scss';

class UploadImages extends React.Component {
  state = {
    previewVisible: false,
    previewImage: '',
    updateList: [],
  }

  componentWillMount() {
    if (this.props.images) {
      const fileList = this.props.images.map(image => (
        {
          uid: image.id,
          name: image.name,
          status: 'done',
          url: image.url,
        }
      ));
      this.setState({ fileList });
    }
  }

  getBase64(img, callback) {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result));
    reader.readAsDataURL(img);
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
          onPreview={this.handlePreview}
          onChange={this.handleChange}
          onRemove={item => this.removeImage(item)}
        >
          {uploadButton}
        </Upload>
        <Modal visible={previewVisible} footer={null} onCancel={this.handleCancel}>
          <img className={styles.modalImage} alt="example" style={{ width: '100%' }} src={previewImage} />
        </Modal>
      </div>
    );
  }
}

export default UploadImages;
