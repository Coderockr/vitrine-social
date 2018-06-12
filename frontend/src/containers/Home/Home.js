import React from 'react';
import Categories from '../../components/Categories';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import api from '../../utils/api';

class Home extends React.Component {
  state = {
    loadingCategories: true,
    loadingRequests: true,
    categories: [],
    requests: [],
    page: 1,
  }

  componentWillMount() {
    this.fetchCategories();
    this.fetchRequests();
  }

  fetchCategories() {
    api.get('categories').then(
      (response) => {
        this.setState({
          categories: response.data,
          loadingCategories: false,
        });
      },
    );
  }

  fetchRequests() {
    api.get(`search?page=${this.state.page}`).then(
      (response) => {
        this.setState({
          requests: response.data,
          loadingRequests: false,
        });
      },
    );
  }

  render() {
    return (
      <Layout>
        <Categories
          loading={this.state.loadingCategories}
          categories={this.state.loadingCategories ? null : this.state.categories}
        />
        <Requests
          loading={this.state.loadingRequests}
          activeRequests={this.state.loadingRequests ? null : this.state.requests}
        />
        <Pagination />
      </Layout>
    );
  }
}

export default Home;
