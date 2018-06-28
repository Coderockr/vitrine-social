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
  AutoComplete,
} from 'antd';
import cx from 'classnames';
import api from '../../utils/api';
import ResponseFeedback from '../ResponseFeedback';
import UploadImages from '../UploadImages';
import { getUser } from '../../utils/auth';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;
const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;

const units = [
  'Itens',
  'Kg',
  'Pessoas',
  'Litros',
];

class RequestForm extends React.Component {
  state = {
    units,
    categories: [],
    responseFeedback: '',
    responseFeedbackMessage: '',
    imagesEnabled: false,
  }

  componentWillMount() {
    this.fetchCategories();
  }

  fetchCategories() {
    api.get('categories').then(
      (response) => {
        this.setState({
          categories: response.data,
        });
      },
    );
  }

  closeModal() {
    this.props.onCancel();
    setTimeout(() => {
      this.props.form.resetFields();
      this.setState({
        responseFeedback: '',
        responseFeedbackMessage: '',
      });
    }, 100);
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        const params = { ...values, organization: getUser().id };
        if (!this.props.request) {
          return api.post('need', params).then(
            () => {
              this.setState({
                responseFeedback: 'success',
                responseFeedbackMessage: 'Nova solicitação adicionada!',
              });
              this.props.onSave();
            },
            () => {
              this.setState({
                responseFeedback: 'error',
                responseFeedbackMessage: 'Não foi possível criar a solicitação!',
              });
            },
          );
        }
        return api.put(`need/${this.props.request.id}`, params).then(
          () => {
            this.setState({
              responseFeedback: 'success',
              responseFeedbackMessage: 'Nova solicitação adicionada!',
            });
            this.props.onSave();
          },
          () => {
            this.setState({
              responseFeedback: 'error',
              responseFeedbackMessage: 'Não foi possível criar a solicitação!',
            });
          },
        );
      }
      return null;
    });
  }

  renderCategories() {
    return (
      this.state.categories.map(category => (
        <Select.Option key={category.id} value={category.id}>{category.name}</Select.Option>
      ))
    );
  }

  renderUnit() {
    return (
      this.state.units.map(unit => (
        <AutoComplete.Option key={unit} value={unit}>{unit}</AutoComplete.Option>
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
        wrapClassName={this.state.responseFeedback && styles.modalFixed}
      >
        <Row className={this.state.responseFeedback && styles.blurRow}>
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
                {getFieldDecorator('status', {
                  initialValue: request ? request.status : 'ACTIVE',
                })(
                  <div className={styles.statusWrapper}>
                    <p className={styles.statusLabel}>Status:</p>
                    <RadioGroup
                      defaultValue={request ? request.status : 'ACTIVE'}
                      className="purpleRadio"
                    >
                      <RadioButton className={styles.radioButton} value="ACTIVE">ATIVA</RadioButton>
                      <RadioButton value="INACTIVE">INATIVA</RadioButton>
                    </RadioGroup>
                  </div>,
                )}
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
                  rules: [{ required: true, message: 'Escolha uma categoria' }],
                  initialValue: request ? request.category.id : '',
                })(
                  <Select
                    placeholder="Categoria"
                    size="large"
                    disabled={request !== null}
                  >
                    {this.renderCategories()}
                  </Select>,
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
                    {getFieldDecorator('requiredQuantity', {
                      rules: [{ required: true, message: 'Preencha a quantidade' }],
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
                    {getFieldDecorator('reachedQuantity', {
                      rules: [{ required: true, message: 'Preencha a quantidade' }],
                      initialValue: request ? request.reachedQuantity : '',
                    })(
                      <InputNumber
                        size="large"
                        min={0}
                        max={Number(this.props.form.getFieldValue('requiredQuantity'))}
                      />,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 12, offset: 0 }}
                  xs={{ span: 8, offset: 2 }}
                >
                  <FormItem label="Tipo">
                    {getFieldDecorator('unit', {
                      rules: [{ required: true, message: 'Escolha um tipo' }],
                      initialValue: request ? request.unit : '',
                    })(
                      <AutoComplete
                        size="large"
                        filterOption
                        disabled={request !== null}
                      >
                        {this.renderUnit()}
                      </AutoComplete>,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              {this.state.imagesEnabled &&
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
              }
              <FormItem>
                <div className={styles.buttonWrapper}>
                  <button
                    className={cx(styles.button, styles.saveButton)}
                    disabled={!this.props.saveEnabled}
                  >
                    SALVAR
                  </button>
                  <button
                    className={cx(styles.button, styles.cancelButton)}
                    onClick={() => this.closeModal()}
                  >
                    CANCELAR
                  </button>
                </div>
              </FormItem>
            </Form>
          </Col>
        </Row>
        <ResponseFeedback
          type={this.state.responseFeedback}
          message={this.state.responseFeedbackMessage}
          onClick={this.state.responseFeedback === 'error' ?
            () => this.setState({ responseFeedback: '', responseFeedbackMessage: '' }) :
            () => this.closeModal()
          }
        />
      </Modal>
    );
  }
}

const WrappedRequestForm = Form.create({
  onValuesChange(props, changedValues, allValues) {
    if (!props.request) {
      return props.enableSave(true);
    }

    let enable = false;
    Object.keys(allValues).forEach((key) => {
      if (key !== 'category' && props.request[key] !== allValues[key]) {
        enable = true;
      }
    });
    return props.enableSave(enable);
  },
})(RequestForm);

export default WrappedRequestForm;
