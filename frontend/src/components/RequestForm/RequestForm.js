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
  Button,
} from 'antd';
import cx from 'classnames';
import ReactGA from 'react-ga';
import { api, apiImage } from '../../utils/api';
import ResponseFeedback from '../ResponseFeedback';
import UploadImages from '../UploadImages';
import { getUser } from '../../utils/auth';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;
const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;

const mediaQuery = window.matchMedia('(max-width: 385px)');

class RequestForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      collapsed: mediaQuery.matches,
      responseFeedback: '',
      responseFeedbackMessage: '',
      imagesChanges: [],
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentDidMount() {
    ReactGA.modalview('/request-form', null, 'Formulário de Solicitação');
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  onChangeImages(imagesChanges) {
    this.setState({ imagesChanges });
    this.props.enableSave(true);
  }

  widthChange() {
    this.setState({
      collapsed: mediaQuery.matches,
    });
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
        const params = {
          ...values,
          organization: getUser().id,
          category: this.props.request.category.id,
        };
        if (params.requiredQuantity === params.reachedQuantity) {
          params.status = 'INACTIVE';
        }
        this.setState({ loading: true });
        Promise.all([
          api.put(`need/${this.props.request.id}`, params),
          ...this.saveImageChanges(),
        ]).then(
          () => {
            this.setState({
              loading: false,
              responseFeedback: 'success',
              responseFeedbackMessage: 'Nova solicitação adicionada!',
            });
            this.props.onSave();
          },
          () => {
            this.setState({
              loading: false,
              responseFeedback: 'error',
              responseFeedbackMessage: 'Não foi possível criar a solicitação!',
            });
          },
        );
      }
    });
  };

  saveImageChanges() {
    const { imagesChanges } = this.state;
    const promisses = [];
    imagesChanges.map((image) => {
      const { file, action } = image;
      if (action === 'add') {
        const formData = new FormData();
        formData.append('file', file);
        formData.append('logo', false);
        return promisses.push(apiImage.post(`need/${this.props.request.id}/images`, formData));
      }
      if (action === 'delete') {
        return promisses.push(apiImage.delete(`need/${this.props.request.id}/images/${file.uid}`));
      }
      return null;
    });
    return promisses;
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
        maskClosable={false}
        wrapClassName={this.state.responseFeedback && styles.modalFixed}
      >
        <Row className={this.state.responseFeedback && styles.blurRow}>
          <Col span={20} offset={2}>
            <h1 className={styles.title}>
              Editar Solicitação
            </h1>
            <Form hideRequiredMark>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('status', {
                  initialValue: request.status,
                })(
                  <div className={styles.statusWrapper}>
                    <p className={styles.statusLabel}>Status:</p>
                    <RadioGroup
                      defaultValue={request.status}
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
                  initialValue: request.title,
                })(
                  <Input size="large" placeholder="Título" disabled={request !== null} />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('category', {
                  initialValue: request.category.name,
                })(
                  <Select
                    placeholder="Categoria"
                    size="large"
                    disabled
                  />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('description', {
                  initialValue: request.description,
                })(
                  <TextArea rows={5} placeholder="Descrição da Solicitação" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={this.state.collapsed ? 10 : 6} className={styles.col}>
                  <FormItem label="Solicitado">
                    {getFieldDecorator('requiredQuantity', {
                      initialValue: request.requiredQuantity,
                    })(
                      <InputNumber size="large" min={1} disabled />,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 6, offset: 0 }}
                  xs={this.state.collapsed ? { span: 10, offset: 4 } : { span: 6, offset: 2 }}
                  className={styles.col}
                >
                  <FormItem label="Recebido">
                    {getFieldDecorator('reachedQuantity', {
                      initialValue: request.reachedQuantity,
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
                  xs={this.state.collapsed ? { span: 24, offset: 0 } : { span: 8, offset: 2 }}
                  className={styles.col}
                >
                  <FormItem label="Tipo">
                    {getFieldDecorator('unit', {
                      initialValue: request.unit,
                    })(
                      <AutoComplete
                        size="large"
                        filterOption
                        disabled
                      />,
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
                  <UploadImages
                    images={request.images}
                    onChange={imagesChanges => this.onChangeImages(imagesChanges)}
                  />
                </Col>
              </FormItem>
              <FormItem>
                <div className={styles.buttonWrapper}>
                  <Button
                    className={cx(styles.button, styles.saveButton)}
                    disabled={!this.props.saveEnabled}
                    onClick={this.handleSubmit}
                    loading={this.state.loading}
                  >
                    SALVAR
                  </Button>
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
  onValuesChange(props, allValues) {
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
