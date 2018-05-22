import React from 'react';
import { Row, Col, Modal, Form, Input, InputNumber, Select } from 'antd';
import cx from 'classnames';
import UploadImages from '../../components/UploadImages';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;
const { Option } = Select;

class RequestDetailsEdit extends React.Component {
  state = {
    types: ['Unidade', 'Kg', 'Pessoa', 'Litro'],
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        return values;
      }
      return null;
    });
  }

  renderType() {
    return (
      this.state.types.map(type => (
        <Option key={type} value={type}>{type}</Option>
      ))
    );
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
            <h1 className={styles.title}>Nova Solicitação</h1>
            <Form
              onSubmit={this.handleSubmit}
              hideRequiredMark
            >
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('title', {
                  rules: [{ required: true, message: 'Preencha o título da solicitação' }],
                })(
                  <Input size="large" placeholder="Título" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('category', {
                  rules: [{ required: true, message: 'Escolha uma Categoria' }],
                })(
                  <Select
                    placeholder="Categoria"
                    size="large"
                    onChange={value => this.getCities(value)}
                  />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <TextArea rows={5} placeholder="Descrição da Solicitação" />
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={6}>
                  <FormItem label="Solicitado">
                    {getFieldDecorator('requestedQty', {
                      rules: [{ required: true, message: 'Preencha o complemento' }],
                    })(
                      <InputNumber size="large" min={1} />,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 6, offset: 0 }}
                  xs={{ span: 6, offset: 2 }}
                >
                  <FormItem label="Recebido">
                    {getFieldDecorator('receivedQty', {
                      rules: [{ required: true, message: 'Preencha o bairro' }],
                    })(
                      <InputNumber size="large" min={1} max={this.props.form.getFieldValue('requestedQty')} />,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 12, offset: 0 }}
                  xs={{ span: 8, offset: 2 }}
                >
                  <FormItem label="Tipo">
                    {getFieldDecorator('type', {
                      rules: [{ required: true, message: 'Escolha um tipo' }],
                    })(
                      <Select size="large">
                        {this.renderType()}
                      </Select>,
                    )}
                  </FormItem>
                </Col>
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
                  <button className={cx(styles.button, styles.saveButton)}>
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

const WrappedEditRequestDetailsForm = Form.create()(RequestDetailsEdit);

export default WrappedEditRequestDetailsForm;
