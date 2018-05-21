import React from 'react';
import { Row, Col, Modal, Avatar, Form, Input, Upload, Select } from 'antd';
import cx from 'classnames';
import UploadImages from '../../components/UploadImages';
import { maskPhone, maskCep } from '../../utils/mask';
import styles from './styles.module.scss';

const FormItem = Form.Item;
const { TextArea } = Input;
const { Option } = Select;

class OrganizationProfileEdit extends React.Component {
  state = {
    validatingZipCode: '',
    states: [{ id: -1, sigla: 'Indisponíveis' }],
    cities: [{ id: -1, nome: 'Indisponíveis' }],
  }

  componentDidMount() {
    this.getStates();
  }

  getCities(stateId) {
    this.props.form.setFieldsValue({
      city: '',
    });
    if (stateId) {
      fetch(`https://servicodados.ibge.gov.br/api/v1/localidades/estados/${stateId}/municipios`)
        .then(response => response.json())
        .then((data) => {
          this.setState({ cities: data.sort(this.citySort) });
        });
    }
  }

  getStates() {
    fetch('https://servicodados.ibge.gov.br/api/v1/localidades/estados')
      .then(response => response.json())
      .then((data) => {
        this.setState({ states: data.sort(this.stateSort) });
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
        const stateId = this.state.states.filter(state => state.sigla === data.uf)[0].id;
        this.getCities(stateId);
        this.props.form.setFieldsValue({
          street: data.logradouro,
          neighborhood: data.bairro,
          state: stateId,
          city: data.localidade,
        });
      })
      .catch(() => {
        this.setState({
          validatingZipCode: 'error',
        });
      });
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
                  getValueFromEvent: e => maskPhone(e.target.value),
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
                validateStatus={this.state.validatingZipCode}
              >
                {getFieldDecorator('zipCode', {
                  getValueFromEvent: e => maskCep(e.target.value),
                  rules: [{
                    required: true, message: 'Preencha o CEP',
                  }, {
                      pattern: /^[0-9]{5}-[0-9]{3}/, message: 'CEP Inválido',
                  }],
                })(
                  <Input
                    ref={(ref) => { this.zipCodeInput = ref; }}
                    size="large"
                    placeholder="CEP"
                    onBlur={() => this.searchZipCode()}
                  />,
                )}
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={17}>
                  <FormItem>
                    {getFieldDecorator('street', {
                      rules: [{ required: true, message: 'Preencha a Rua' }],
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
                    })(
                      <Input size="large" placeholder="Número" />,
                    )}
                  </FormItem>
                </Col>
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Col span={9}>
                  <FormItem>
                    {getFieldDecorator('complement', {
                      rules: [{ required: true, message: 'Preencha o complemento' }],
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
                    })(
                      <Input size="large" placeholder="Bairro" />,
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
                    })(
                      <Select placeholder="Cidade" size="large" showSearch optionFilterProp="children">
                        {this.renderCities()}
                      </Select>,
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

const WrappedEditProfileForm = Form.create()(OrganizationProfileEdit);

export default WrappedEditProfileForm;
