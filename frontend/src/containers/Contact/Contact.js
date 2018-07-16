import React from 'react';
import { Row, Col, Form, Input, Select } from 'antd';
import cx from 'classnames';
import ReactGA from 'react-ga';
import styles from './styles.module.scss';
import Layout from '../../components/Layout';
import BottomNotification from '../../components/BottomNotification';
import { maskPhone } from '../../utils/mask';
import api from '../../utils/api';

const FormItem = Form.Item;
const { TextArea } = Input;

const reasons = [
  'Dúvida',
  'Cadastro de Entidade',
  'Sugestão',
  'Reclamação',
  'Outros',
];

class Contact extends React.Component {
  componentDidMount() {
    document.title = 'Vitrine Social - Contato';
  }

  handleSubmit = (e) => {
    ReactGA.event({
      category: 'Usuario',
      action: 'Enviar Contato Coderockr',
    });
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        api.post('contact', values).then(
          () => BottomNotification({ message: 'Formulário de contato enviado!', success: true }),
          () => BottomNotification({ message: 'Não foi possível enviar o formulário!', success: false }),
        );
      }
    });
  }

  renderReasons() {
    return (
      reasons.map(reason => (
        <Select.Option key={reason} value={reason}>{reason}</Select.Option>
      ))
    );
  }

  render() {
    const { getFieldDecorator } = this.props.form;
    const formItemLayout = {
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 14, offset: 5 },
        lg: { span: 16, offset: 4 },
      },
    };

    return (
      <Layout>
        <Row>
          <Col span={20} offset={2}>
            <div className={styles.contactWrapper}>
              <h1 className={styles.title}>ENTRE EM CONTATO</h1>
              <Form onSubmit={this.handleSubmit}>
                <FormItem
                  {...formItemLayout}
                >
                  {getFieldDecorator('name', {
                    rules: [{ required: true, message: 'Preencha com o seu nome' }],
                  })(
                    <Input size="large" placeholder="Nome" />,
                  )}
                </FormItem>
                <FormItem
                  {...formItemLayout}
                >
                  {getFieldDecorator('email', {
                    rules: [{
                      type: 'email', message: 'E-mail inválido',
                    }, {
                      required: true, message: 'Preencha com o seu e-mail',
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
                      pattern: /^1\d\d(\d\d)?$|^0800 ?\d{3} ?\d{4}$|^(\(0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d\) ?|0?([1-9a-zA-Z][0-9a-zA-Z])?[1-9]\d[ .-]?)?(9|9[ .-])?[2-9]\d{3}[ .-]?\d{4}$/gm, message: 'Telefone Inválido',
                    }],
                  })(
                    <Input size="large" placeholder="Telefone" />,
                  )}
                </FormItem>
                <FormItem
                  {...formItemLayout}
                >
                  {getFieldDecorator('reason', {
                    rules: [{
                      required: true, message: 'Escolha o motivo do seu contato',
                    }],
                  })(
                    <Select
                      placeholder="Motivo do contato"
                      size="large"
                    >
                      {this.renderReasons()}
                    </Select>,
                  )}
                </FormItem>
                <FormItem
                  {...formItemLayout}
                >
                  {getFieldDecorator('message', {
                    rules: [{
                      required: true, message: 'Escreva uma mensagem',
                    }],
                  })(
                    <TextArea rows={5} placeholder="Mensagem" />,
                  )}
                </FormItem>
                <FormItem>
                  <button className={cx(styles.button, styles.sendButton)} htmlType="submit">
                    ENVIAR
                  </button>
                </FormItem>
              </Form>
            </div>
          </Col>
        </Row>
      </Layout>
    );
  }
}

const WrappedContactForm = Form.create()(Contact);

export default WrappedContactForm;
