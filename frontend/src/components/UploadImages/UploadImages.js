import React from 'react';
import { Upload, Icon, Modal } from 'antd';
import styles from './styles.module.scss';

class UploadImages extends React.Component {
  state = {
    previewVisible: false,
    previewImage: '',
  }

  componentWillMount() {
    const fileList = this.props.images.map(image => (
      {
        uid: image.id,
        name: image.name.images,
        status: 'done',
        url: image.url,
      }
    ));
    this.setState({ fileList });
  }

  handleCancel = () => this.setState({ previewVisible: false })

  handlePreview = (file) => {
    this.setState({
      previewImage: file.url || file.thumbUrl,
      previewVisible: true,
    });
  }

  handleChange = ({ fileList }) => this.setState({ fileList })

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
          action="//jsonplaceholder.typicode.com/posts/"
          listType="picture-card"
          fileList={fileList}
          onPreview={this.handlePreview}
          onChange={this.handleChange}
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
