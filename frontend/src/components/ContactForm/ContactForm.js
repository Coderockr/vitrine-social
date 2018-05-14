import React from 'react';
import moment from 'moment';
import { Row, Col, Form, Input } from 'antd';
import styles from './styles.module.scss';

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

  render() {
    const formItemLayout = {
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 14, offset: 5 },
        lg: { span: 16, offset: 4 },
      },
    };

    return (
      <Row>
        <Col span={20} offset={2}>
          <div className={styles.contactWrapper}>
            <p>Entre em contato com</p>
            <p className={styles.organizationName}>{request.organization.name}</p>
            <p>Telefone: <span>{request.organization.phoneNumber}</span></p>
            <p className={styles.fillLabel}>Ou preencha o formulário:</p>
            <Form>
              <FormItem
                {...formItemLayout}
              >
                <Input size="large" placeholder="Nome" />
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Input size="large" placeholder="E-mail" />
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <Input size="large" placeholder="Telefone" />
              </FormItem>
              <FormItem
                {...formItemLayout}
              >
                <TextArea rows={5} placeholder="Mensagem" />
              </FormItem>
            </Form>
            <button className={styles.button}>ENVIAR</button>
          </div>
        </Col>
      </Row>
    );
  }
}

export default ContactForm;
