import React from 'react';
import moment from 'moment';
import { Row, Col, Form, Input } from 'antd';
import Icon from '../../components/Icons';
import styles from './styles.module.scss';
import colors from '../../utils/styles/colors';

const FormItem = Form.Item;
const { TextArea } = Input;

const request = {
  organization: {
    name: 'Lar Abdon Batista',
    link: 'http://coderockr.com/',
    phoneNumber: '(47) 3227-6359',
  },
  category: 'voluntarios',
  data: moment().subtract(2, 'days'),
  item: '10 voluntários para ler para criancinhas felizes',
  description: 'v-governmental organizations, nongovernmental organizations, or nongovernment organizations, commonly referred to as NGOs, are nonprofit organizations independent of governments and international',
};

class ContactForm extends React.Component {
  state = {
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
      <Row>
        <Col span={24}>
          <button className={styles.backButton} onClick={this.props.onClick}>
            <Icon icon="reply" size={40} color={colors.white} />
          </button>
        </Col>
        <Col span={20} offset={2}>
          <div className={styles.contactWrapper}>
            <p>Entre em contato com</p>
            <p className={styles.organizationName}>{request.organization.name}</p>
            <p>Telefone: <span>{request.organization.phoneNumber}</span></p>
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
                <TextArea rows={5} placeholder="Mensagem" />
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
    );
  }
}

const WrappedContactForm = Form.create()(ContactForm);

export default WrappedContactForm;
