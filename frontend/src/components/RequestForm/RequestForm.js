import React from 'react';
import {
  Row,
  Col,
  Modal,
  Form,
  Input,
  InputNumber,
  Select,
  Radio,
} from 'antd';
import cx from 'classnames';
import UploadImages from '../UploadImages';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;
const { Option } = Select;
const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;

class RequestForm extends React.Component {
  state = {
    types: [
      'Unidade',
      'Kg',
      'Pessoa',
      'Litro',
    ],
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
    const { request } = this.props;

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
            <h1 className={styles.title}>
              {request ? 'Editar Solicitação' : 'Nova Solicitação'}
            </h1>
            <Form
              onSubmit={this.handleSubmit}
              hideRequiredMark
            >
              <FormItem
                className={request ? null : styles.statusFormItem}
                {...formItemLayout}
              >
                <div className={styles.statusWrapper}>
                  <p className={styles.statusLabel}>Status:</p>
                  <RadioGroup defaultValue={request ? request.status : 'ACTIVE'} className="purpleRadio">
                    <RadioButton className={styles.radioButton} value="ACTIVE">ATIVA</RadioButton>
                    <RadioButton value="INACTIVE">INATIVA</RadioButton>
                  </RadioGroup>
                </div>
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('title', {
                  rules: [{ required: true, message: 'Preencha o título da solicitação' }],
                  initialValue: request ? request.title : '',
                })(
                  <Input size="large" placeholder="Título" disabled={request !== null} />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('category', {
                  rules: [{ required: true, message: 'Escolha uma Categoria' }],
                  initialValue: request ? request.category.name : '',
                })(
                  <Select
                    placeholder="Categoria"
                    size="large"
                    disabled={request !== null}
                  />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('description', {
                  rules: [{ required: true, message: 'Descreva a solicitação' }],
                  initialValue: request ? request.description : '',
                })(
                  <TextArea rows={5} placeholder="Descrição da Solicitação" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={6}>
                  <FormItem label="Solicitado">
                    {getFieldDecorator('requestedQty', {
                      rules: [{ required: true, message: 'Preencha a quantidade solicitada' }],
                      initialValue: request ? request.requiredQuantity : '',
                    })(
                      <InputNumber size="large" min={1} disabled={request !== null} />,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 6, offset: 0 }}
                  xs={{ span: 6, offset: 2 }}
                >
                  <FormItem label="Recebido">
                    {getFieldDecorator('receivedQty', {
                      rules: [{ required: true, message: 'Preencha a quantidade recebida' }],
                      initialValue: request ? request.reachedQuantity : '',
                    })(
                      <InputNumber
                        size="large"
                        min={1}
                        max={this.props.form.getFieldValue('requestedQty')}
                      />,
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
                      initialValue: request ? request.unit : '',
                    })(
                      <Select size="large" disabled={request !== null}>
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
                  <UploadImages images={request ? request.images : null} />
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

const WrappedRequestForm = Form.create()(RequestForm);

export default WrappedRequestForm;
