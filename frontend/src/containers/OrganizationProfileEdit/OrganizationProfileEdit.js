import React from 'react';
import { Row, Col, Modal, Avatar, Form, Input, Upload } from 'antd';
import cx from 'classnames';
import UploadImages from '../../components/UploadImages';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;

class OrganizationProfileEdit extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        return values;
      }
      return null;
    });
  }

  render() {
    const { getFieldDecorator } = this.props.form;
    const formItemLayout = {
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 20, offset: 2 },
        md: { span: 16, offset: 4 },
      },
    };

    return (
      <Modal
        visible={this.props.visible}
        footer={null}
        width={800}
        className={styles.modal}
        destroyOnClose
        onCancel={this.props.onCancel}
        closable={false}
        maskClosable={false}
      >
        <Row>
          <Col span={20} offset={2}>
            <h1 className={styles.title}>Editar Perfil da Organização</h1>
            <div className={styles.uploadWrapper}>
              <Upload
                name="avatar"
                listType="picture"
                showUploadList={false}
                action="//jsonplaceholder.typicode.com/posts/"
                onChange={this.handleChange}
              >
                <div className={styles.avatarWrapper}>
                  <Avatar
                    icon="user"
                    style={{
                      fontSize: 140,
                      color: '#FFFFFF',
                      backgroundColor: '#FFE7D5',
                      textShadow: '4px 1px 3px #FF974A',
                    }}
                  />
                  <div className={styles.avatarEdit}>
                    <p>Editar</p>
                  </div>
                </div>
              </Upload>
            </div>
            <Form onSubmit={this.handleSubmit}>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('name', {
                  rules: [{ required: true, message: 'Preencha o nome da organização' }],
                })(
                  <Input size="large" placeholder="Nome da Organização" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('email', {
                  rules: [{
                    type: 'email', message: 'E-mail inválido',
                  }, {
                    required: true, message: 'Preencha o e-mail',
                  }],
                })(
                  <Input size="large" placeholder="E-mail" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('phone', {
                  rules: [{
                    required: true, message: 'Preencha o telefone',
                  }, {
                    pattern: /^1\d\d(\d\d)?$|^0800 ?\d{3} ?\d{4}$|^(\(0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d\) ?|0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d[ .-]?)?(9|9[ .-])?[2-9]\d{3}[ .-]?\d{4}$/gm, message: 'Telefone Inválido',
                  }],
                })(
                  <Input size="large" placeholder="Telefone" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('zipCode', {
                  rules: [{ required: true, message: 'Preencha o CEP' }],
                })(
                  <Input size="large" placeholder="CEP" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('street', {
                  rules: [{ required: true, message: 'Preencha a Rua' }],
                })(
                  <Input size="large" placeholder="Rua" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={9}>
                  <FormItem>
                    {getFieldDecorator('number', {
                      rules: [{ required: true, message: 'Preencha o número' }],
                    })(
                      <Input size="large" placeholder="Número" />,
                    )}
                  </FormItem>
                </Col>
                <Col span={1}>
                  <span style={{ display: 'inline-block', width: '100%', textAlign: 'center' }} />
                </Col>
                <Col span={14}>
                  <FormItem>
                    {getFieldDecorator('complement', {
                      rules: [{ required: true, message: 'Preencha o complemento' }],
                    })(
                      <Input size="large" placeholder="Complemento" />,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={9}>
                  <FormItem>
                    {getFieldDecorator('state', {
                      rules: [{ required: true, message: 'Preencha o Estado' }],
                    })(
                      <Input size="large" placeholder="Estado" />,
                    )}
                  </FormItem>
                </Col>
                <Col span={1}>
                  <span style={{ display: 'inline-block', width: '100%', textAlign: 'center' }} />
                </Col>
                <Col span={14}>
                  <FormItem>
                    {getFieldDecorator('city', {
                      rules: [{ required: true, message: 'Preencha a Cidade' }],
                    })(
                      <Input size="large" placeholder="Cidade" />,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <TextArea rows={5} placeholder="Sobre a Organização" />
              </FormItem>
              <FormItem>
                <Col
                  md={{ span: 18, offset: 3 }}
                  sm={{ span: 22, offset: 1 }}
                  xs={{ span: 24, offset: 0 }}
                >
                  <h2 className={styles.galleryHeader}>Galeria de Imagens</h2>
                  <UploadImages />
                </Col>
              </FormItem>
              <FormItem>
                <div className={styles.buttonWrapper}>
                  <button className={cx(styles.button, styles.saveButton)} htmlType="submit">
                    SALVAR
                  </button>
                  <button
                    className={cx(styles.button, styles.cancelButton)}
                    onClick={this.props.onCancel}
                  >
                    CANCELAR
                  </button>
                </div>
              </FormItem>
            </Form>
          </Col>
        </Row>
      </Modal>
    );
  }
}

const WrappedEditProfileForm = Form.create()(OrganizationProfileEdit);

export default WrappedEditProfileForm;
