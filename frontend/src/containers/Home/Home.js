import React from 'react';
import Search from '../../components/Search';
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
          requests: response.data.results,
          loadingRequests: false,
        });
      },
    );
  }

  searchRequests(text) {
    const { history } = this.props;
    history.push(`/search/text=${text}&page=1`);
  }

  render() {
    return (
      <Layout>
        <Search search={text => this.searchRequests(text)} />
        <Categories
          loading={this.state.loadingCategories}
          categories={this.state.loadingCategories ? null : this.state.categories}
          hasSearch
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
