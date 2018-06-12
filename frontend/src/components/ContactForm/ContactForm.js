import React from 'react';
import { Row, Col, Form, Input } from 'antd';
import ResponseFeedback from '../ResponseFeedback';
import Icon from '../../components/Icons';
import styles from './styles.module.scss';
import colors from '../../utils/styles/colors';
import { maskPhone } from '../../utils/mask';
import api from '../../utils/api';

const FormItem = Form.Item;
const { TextArea } = Input;

class ContactForm extends React.Component {
  state = {
    responseFeedback: '',
    responseFeedbackMessage: '',
  }

  closeModal() {
    this.props.onFeedback(false);
    this.props.onClick();
    setTimeout(() => {
      this.props.form.resetFields();
      this.setState({
        responseFeedback: '',
        responseFeedbackMessage: '',
      });
    }, 100);
  }

  backToForm() {
    this.props.onFeedback(false);
    this.setState({
      responseFeedback: '',
      responseFeedbackMessage: '',
    });
  }

  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        return api.post(`need/${this.props.request.id}/response`, values).then(
          () => {
            this.props.onFeedback(true);
            this.setState({
              responseFeedback: 'success',
              responseFeedbackMessage: 'Formulário de contato enviado!',
            });
            this.props.onSave();
          },
          () => {
            this.props.onFeedback(true);
            this.setState({
              responseFeedback: 'error',
              responseFeedbackMessage: 'Não foi possível enviar o formulário!',
            });
          },
        );
      }
      return null;
    });
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

    const { request } = this.props;

    return (
      <div>
        <Row className={this.state.responseFeedback && styles.blurRow}>
          <Col span={24}>
            <button className={styles.backButton} onClick={this.props.onClick}>
              <Icon icon="reply" size={40} color={colors.white} />
            </button>
          </Col>
          <Col span={20} offset={2}>
            <div className={styles.contactWrapper}>
              <p>Entre em contato com</p>
              <p className={styles.organizationName}>{request.organization.name}</p>
              <p>Telefone: <span>{request.organization.phone}</span></p>
              <p className={styles.fillLabel}>Ou preencha o formulário:</p>
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
                      required: true, message: 'Preencha com o seu telefone',
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
                  {getFieldDecorator('message', {
                    rules: [{
                      required: true, message: 'Escreva uma mensagem',
                    }],
                  })(
                    <TextArea rows={5} placeholder="Mensagem" />,
                  )}
                </FormItem>
                <FormItem>
                  <button htmlType="submit">
                    ENVIAR
                  </button>
                </FormItem>
              </Form>
            </div>
          </Col>
        </Row>
        <ResponseFeedback
          type={this.state.responseFeedback}
          message={this.state.responseFeedbackMessage}
          onClick={this.state.responseFeedback === 'error' ?
            () => this.backToForm() :
            () => this.closeModal()
          }
        />
      </div>
    );
  }
}

const WrappedContactForm = Form.create()(ContactForm);

export default WrappedContactForm;
