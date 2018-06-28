import React from 'react';
import {
  Row,
  Col,
  Modal,
  Avatar,
  Form,
  Input,
  Upload,
  Select,
} from 'antd';
import cx from 'classnames';
import UploadImages from '../UploadImages';
import ResponseFeedback from '../ResponseFeedback';
import { maskPhone, maskCep } from '../../utils/mask';
import api from '../../utils/api';
import { getUser } from '../../utils/auth';
import colors from '../../utils/styles/colors';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;
const { Option } = Select;

class OrganizationProfileForm extends React.Component {
  state = {
    validatingZipCode: '',
    responseFeedback: '',
    responseFeedbackMessage: '',
    states: [{ id: -1, sigla: 'Indisponíveis' }],
    cities: [{ id: -1, nome: 'Indisponíveis' }],
    imagesEnabled: false,
  }

  componentDidMount() {
    this.getStates();
  }

  getStateId(initials) {
    return this.state.states.filter(state => state.sigla === initials)[0].id;
  }

  getStateName(stateId) {
    return this.state.states.filter(state => state.id === stateId)[0].sigla;
  }

  getCities(stateId, selectedCity) {
    this.props.form.setFieldsValue({
      city: '',
    });
    if (stateId) {
      fetch(`https://servicodados.ibge.gov.br/api/v1/localidades/estados/${stateId}/municipios`)
        .then(response => response.json())
        .then((data) => {
          this.setState({ cities: data.sort(this.citySort) });
          if (selectedCity) {
            this.props.form.setFieldsValue({
              city: selectedCity,
            });
          }
        });
    }
  }

  getStates() {
    fetch('https://servicodados.ibge.gov.br/api/v1/localidades/estados')
      .then(response => response.json())
      .then((data) => {
        const states = data.sort(this.stateSort);
        this.setState({ states });
        const { organization } = this.props;
        if (organization) {
          const stateId = this.getStateId(organization.address.state);
          this.getCities(stateId, organization.address.city);
          this.props.form.setFieldsValue({
            state: stateId,
          });
        }
      });
  }

  stateSort(a, b) {
    if (a.sigla < b.sigla) {
      return -1;
    }
    if (a.sigla > b.sigla) {
      return 1;
    }
    return 0;
  }

  citySort(a, b) {
    if (a.nome < b.nome) {
      return -1;
    }
    if (a.nome > b.nome) {
      return 1;
    }
    return 0;
  }

  searchZipCode() {
    this.setState({ validatingZipCode: 'validating' });
    fetch(`https://viacep.com.br/ws/${this.zipCodeInput.props.value}/json/ `)
      .then(response => response.json())
      .then((data) => {
        this.setState({
          validatingZipCode: 'success',
        });
        const stateId = this.getStateId(data.uf);
        this.getCities(stateId, data.localidade);
        this.props.form.setFieldsValue({
          street: data.logradouro,
          neighborhood: data.bairro,
          state: stateId,
        });
      })
      .catch(() => {
        this.setState({
          validatingZipCode: 'error',
        });
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
        const address = {
          zipcode: values.zipcode,
          street: values.street,
          number: values.number,
          complement: values.complement,
          neighborhood: values.neighborhood,
          state: this.getStateName(values.state),
          city: values.city,
        };
        const params = {
          name: values.name,
          email: values.email,
          phone: values.phone,
          address,
          about: values.about,
        };
        api.put(`organization/${getUser().id}`, params).then(
          () => {
            this.setState({
              responseFeedback: 'success',
              responseFeedbackMessage: 'Dados da organização salvos!',
            });
            this.props.onSave();
          },
          () => {
            this.setState({
              responseFeedback: 'error',
              responseFeedbackMessage: 'Não foi possível salvar os dados da organização!',
            });
          },
        );
      }
    });
  }

  renderStates() {
    return (
      this.state.states.map(state => (
        <Option key={state.id} value={state.id}>{state.sigla}</Option>
      ))
    );
  }

  renderCities() {
    return (
      this.state.cities.map(city => (
        <Option key={city.id} value={city.id}>{city.nome}</Option>
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

    const { organization } = this.props;
    const { address } = organization;

    return (
      <Modal
        visible={this.props.visible}
        footer={null}
        width={800}
        className={styles.modal}
        destroyOnClose
        closable={false}
        maskClosable={false}
        wrapClassName={this.state.responseFeedback && styles.modalFixed}
      >
        <Row className={this.state.responseFeedback && styles.blurRow}>
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
                      color: colors.white,
                      backgroundColor: colors.ambar_200,
                      textShadow: '4px 1px 3px #FF974A',
                    }}
                    src={organization.logo}
                  />
                  <div className={styles.avatarEdit}>
                    <p>Editar</p>
                  </div>
                </div>
              </Upload>
            </div>
            <Form onSubmit={this.handleSubmit}>
              <FormItem {...formItemLayout}>
                {getFieldDecorator('name', {
                  rules: [{ required: true, message: 'Preencha o nome da organização' }],
                  initialValue: organization.name,
                })(
                  <Input size="large" placeholder="Nome da Organização" />,
                )}
              </FormItem>
              <FormItem {...formItemLayout}>
                {getFieldDecorator('email', {
                  rules: [{
                    type: 'email', message: 'E-mail inválido',
                  }, {
                    required: true, message: 'Preencha o e-mail',
                  }],
                  initialValue: organization.email,
                })(
                  <Input size="large" placeholder="E-mail" />,
                )}
              </FormItem>
              <FormItem {...formItemLayout}>
                {getFieldDecorator('phone', {
                  getValueFromEvent: e => maskPhone(e.target.value),
                  rules: [{
                    required: true, message: 'Preencha o telefone',
                  }, {
                    pattern: /^1\d\d(\d\d)?$|^0800 ?\d{3} ?\d{4}$|^(\(0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d\) ?|0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d[ .-]?)?(9|9[ .-])?[2-9]\d{3}[ .-]?\d{4}$/gm, message: 'Telefone Inválido',
                  }],
                  initialValue: maskPhone(organization.phone),
                })(
                  <Input size="large" placeholder="Telefone" />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
                validateStatus={this.state.validatingZipCode}
              >
                {getFieldDecorator('zipcode', {
                  getValueFromEvent: e => maskCep(e.target.value),
                  rules: [{
                    required: true, message: 'Preencha o CEP',
                  }, {
                      pattern: /^[0-9]{5}-[0-9]{3}/, message: 'CEP Inválido',
                  }],
                  initialValue: maskCep(address.zipcode),
                })(
                  <Input
                    ref={(ref) => { this.zipCodeInput = ref; }}
                    size="large"
                    placeholder="CEP"
                    onBlur={() => this.searchZipCode()}
                  />,
                )}
              </FormItem>
              <FormItem {...formItemLayout}>
                <Col span={17}>
                  <FormItem>
                    {getFieldDecorator('street', {
                      rules: [{ required: true, message: 'Preencha a Rua' }],
                      initialValue: address.street,
                    })(
                      <Input size="large" placeholder="Rua" />,
                    )}
                  </FormItem>
                </Col>
                <Col span={1}>
                  <span style={{ display: 'inline-block', width: '100%', textAlign: 'center' }} />
                </Col>
                <Col span={6}>
                  <FormItem>
                    {getFieldDecorator('number', {
                      rules: [{ required: true, message: 'Preencha o número' }],
                      initialValue: address.number,
                    })(
                      <Input size="large" placeholder="Número" />,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem {...formItemLayout}>
                <Col span={9}>
                  <FormItem>
                    {getFieldDecorator('complement', {
                      rules: [{ required: true, message: 'Preencha o complemento' }],
                      initialValue: address.complement,
                    })(
                      <Input size="large" placeholder="Complemento" />,
                    )}
                  </FormItem>
                </Col>
                <Col span={1}>
                  <span style={{ display: 'inline-block', width: '100%', textAlign: 'center' }} />
                </Col>
                <Col span={14}>
                  <FormItem>
                    {getFieldDecorator('neighborhood', {
                      rules: [{ required: true, message: 'Preencha o bairro' }],
                      initialValue: address.neighborhood,
                    })(
                      <Input size="large" placeholder="Bairro" />,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem {...formItemLayout}>
                <Col span={9}>
                  <FormItem>
                    {getFieldDecorator('state', {
                      rules: [{ required: true, message: 'Preencha o Estado' }],
                      initialValue: this.state.states.length > 1 ?
                        this.getStateId(address.state) :
                        null,
                    })(
                      <Select
                        placeholder="Estado"
                        size="large"
                        onChange={value => this.getCities(value)}
                      >
                        {this.renderStates()}
                      </Select>,
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
                      initialValue: address.city,
                    })(
                      <Select placeholder="Cidade" size="large" showSearch optionFilterProp="children">
                        {this.renderCities()}
                      </Select>,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem {...formItemLayout}>
                {getFieldDecorator('about', { initialValue: organization.about })(
                  <TextArea rows={5} placeholder="Sobre a Organização" />,
                )}
              </FormItem>
              {this.state.imagesEnabled &&
                <FormItem>
                  <Col
                    md={{ span: 18, offset: 3 }}
                    sm={{ span: 22, offset: 1 }}
                    xs={{ span: 24, offset: 0 }}
                  >
                    <h2 className={styles.galleryHeader}>Galeria de Imagens</h2>
                    <UploadImages images={organization.images} />
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

const WrappedOrganizationProfileForm = Form.create({
  onValuesChange(props, changedValues, allValues) {
    const allCurrentValues = { ...props.organization, ...props.organization.address };
    const { city } = props.organization.address;
    if (!changedValues.state && changedValues.city !== '' && changedValues.city !== city) {
      let enable = false;
      Object.keys(allValues).forEach((key) => {
        if (key !== 'state' && allCurrentValues[key] !== allValues[key]) {
          enable = true;
        }
      });
      props.enableSave(enable);
    }
  },
})(OrganizationProfileForm);

export default WrappedOrganizationProfileForm;
