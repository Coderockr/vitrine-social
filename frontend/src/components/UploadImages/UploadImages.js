import React from 'react';
import { Upload, Icon, Modal } from 'antd';
import api from '../../utils/api';
import { getUser } from '../../utils/auth';
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
    let updateFile = { file: item, action: 'delete' };
    let updateList = [...this.state.updateList];
    if (find) {
      updateFile = { ...find };
      updateFile.action = 'delete';
      const index = updateList.indexOf(find);
      updateList.splice(index, 1);
    }

    updateList = [...updateList, updateFile];
    this.setState({ updateList });
    this.props.onChange(updateList);
  }

  uploadImage({ onSuccess, onError, file }) {
    const formData = new FormData();
    formData.append('images', file);
    api.post(`organization/${getUser().id}/images`, formData).then((result) => {
      const updateFile = { file, action: 'add', uid: result.data.id };
      const updateList = [...this.state.updateList, updateFile];
      this.setState({ updateList });
      this.props.onChange(updateList);
      onSuccess(null, updateFile);
    }).catch(() => (
      onError()
    ));
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
