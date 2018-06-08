import React from 'react';
import Categories from '../../components/Categories';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import api from '../../utils/api';

class Home extends React.Component {
  state = {
    loading: true,
    categories: [],
    requests: [],
  }

  componentWillMount() {
    this.fetchCategories();
  }

  fetchCategories() {
    api.get('categories').then(
      (response) => {
        this.setState({
          categories: response.data,
          loading: false,
        });
      },
    );
  }

  render() {
    return (
      <Layout>
        <Categories
          loading={this.state.loading}
          categories={this.state.loading ? null : this.state.categories}
        />
        <Requests
          loading={this.state.loading}
          activeRequests={this.state.loading ? null : this.state.requests}
        />
        <Pagination />
      </Layout>
    );
  }
}

export default Home;
