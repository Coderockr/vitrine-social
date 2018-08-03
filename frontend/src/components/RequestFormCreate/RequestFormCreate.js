import React from 'react';
import {
  Row,
  Col,
  Modal,
  Form,
  Input,
  InputNumber,
  Select,
  AutoComplete,
  Button,
} from 'antd';
import cx from 'classnames';
import ReactGA from 'react-ga';
import api from '../../utils/api';
import ResponseFeedback from '../ResponseFeedback';
import UploadImages from '../UploadImages';
import { getUser } from '../../utils/auth';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;

const mediaQuery = window.matchMedia('(max-width: 385px)');

const units = [
  'kg',
  'litros',
  'caixas',
  'peças',
  'pessoas',
  'unidades',
];

class RequestForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      collapsed: mediaQuery.matches,
      units,
      categories: [],
      responseFeedback: '',
      responseFeedbackMessage: '',
      imagesChanges: [],
    };

    mediaQuery.addListener(this.widthChange.bind(this));
  }

  componentWillMount() {
    ReactGA.modalview('/request-form', null, 'Formulário de Nova Solicitação');
    this.fetchCategories();
  }

  componentWillUnmount() {
    mediaQuery.removeListener(this.widthChange);
  }

  widthChange() {
    this.setState({
      collapsed: mediaQuery.matches,
    });
  }

  fetchCategories() {
    api.get('categories').then(
      (response) => {
        this.setState({ categories: response.data });
      },
    );
  }

  resetForm() {
    this.props.form.resetFields();
    this.setState({
      responseFeedback: '',
      responseFeedbackMessage: '',
    });
  }

  closeModal() {
    this.props.onCancel();
    setTimeout(() => {
      this.resetForm();
    }, 100);
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        this.setState({ loading: true });
        const params = { ...values, organization: getUser().id };
        api.post('need', params).then(result => this.saveImageChanges(result)
          .then(
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
          ),
        );
      }
    });
  }

  saveImageChanges(result) {
    const { imagesChanges } = this.state;
    const promisses = [];
    imagesChanges.map((image) => {
      const { file } = image;
      const formData = new FormData();
      formData.append('file', file);
      formData.append('logo', false);
      return promisses.push(api.post(`need/${result.data.id}/images`, formData));
    });
    return Promise.all(promisses);
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

  renderAdditionalButtons() {
    return (
      <button
        className={cx(styles.button, styles.additionalButton)}
        onClick={() => this.resetForm()}
      >
        ADICIONAR OUTRA
      </button>
    );
  }

  render() {
    const { getFieldDecorator } = this.props.form;

    const formItemLayout = {
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 22, offset: 1 },
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
            <h1 className={styles.title}>Nova Solicitação</h1>
            <Form hideRequiredMark>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('category', {
                  rules: [{ required: true, message: 'Escolha uma categoria' }],
                })(
                  <Select
                    placeholder="Categoria"
                    size="large"
                  >
                    {this.renderCategories()}
                  </Select>,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={this.state.collapsed ? 8 : 6} className={styles.col}>
                  <FormItem label="Quantidade">
                    {getFieldDecorator('requiredQuantity', {
                      rules: [{ required: true, message: 'Obrigatório' }],
                    })(
                      <InputNumber size="large" min={1} />,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 7, offset: 0 }}
                  xs={this.state.collapsed ? { span: 14, offset: 2 } : { span: 16, offset: 2 }}
                  className={styles.col}
                >
                  <FormItem label="Tipo">
                    {getFieldDecorator('unit', {
                      rules: [{ required: true, message: 'Ex: kg, litros' }],
                    })(
                      <AutoComplete
                        size="large"
                        filterOption
                      >
                        {this.renderUnit()}
                      </AutoComplete>,
                    )}
                  </FormItem>
                </Col>
                <Col
                  sm={{ span: 10, offset: 1 }}
                  xs={{ span: 24, offset: 0 }}
                  className={styles.col}
                >
                  <FormItem label="Item">
                    {getFieldDecorator('title', {
                      rules: [{ required: true, message: 'O que você precisa?' }],
                    })(
                      <Input size="large" />,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                {getFieldDecorator('description')(
                  <TextArea rows={5} placeholder="Descrição da Solicitação" />,
                )}
              </FormItem>
              <FormItem>
                <Col
                  md={{ span: 18, offset: 3 }}
                  sm={{ span: 22, offset: 1 }}
                  xs={{ span: 24, offset: 0 }}
                >
                  <h2 className={styles.galleryHeader}>Galeria de Imagens</h2>
                  <UploadImages onChange={imagesChanges => this.setState({ imagesChanges })} />
                </Col>
              </FormItem>
              <FormItem>
                <div className={styles.buttonWrapper}>
                  <Button
                    className={cx(styles.button, styles.saveButton)}
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
          additionalButtons={this.state.responseFeedback === 'error' ? null : () => this.renderAdditionalButtons()}
        />
      </Modal>
    );
  }
}

const WrappedRequestForm = Form.create()(RequestForm);

export default WrappedRequestForm;
