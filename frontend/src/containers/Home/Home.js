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

  componentDidMount() {
    document.title = 'Vitrine Social';
  }

  onChangePage(page) {
    this.setState({
      page,
      loadingRequests: true,
    }, () => {
      this.fetchRequests();
    });
  }

  fetchCategories() {
    api.get('categories').then(
      (response) => {
        this.setState({
          categories: response.data,
          loadingCategories: false,
        });
      }, (error) => {
        this.setState({
          loadingCategories: false,
          errorCategories: error,
        });
      },
    );
  }

  fetchRequests() {
    api.get(`search?page=${this.state.page}&status=ACTIVE`).then(
      (response) => {
        this.setState({
          requests: response.data.results,
          pagination: response.data.pagination,
          loadingRequests: false,
        });
      }, (error) => {
        this.setState({
          loadingRequests: false,
          errorRequests: error,
        });
      },
    );
  }

  searchRequests(text) {
    const { history } = this.props;
    history.push(`/search/text=${text}&page=1&status=ACTIVE`);
  }

  searchByCategory(categoryId) {
    const { history } = this.props;
    history.push(`/search/categories=${categoryId}&page=1&status=ACTIVE`);
  }

  render() {
    return (
      <Layout>
        <Search search={text => this.searchRequests(text)} />
        <Categories
          loading={this.state.loadingCategories}
          categories={this.state.loadingCategories ? null : this.state.categories}
          hasSearch
          onClick={id => this.searchByCategory(id)}
          error={this.state.errorCategories}
        />
        <Requests
          loading={this.state.loadingRequests}
          activeRequests={this.state.loadingRequests ? null : this.state.requests}
          error={this.state.errorRequests}
        />
        {this.state.pagination &&
        <Pagination
          current={this.state.page}
          total={this.state.pagination.totalResults}
          onChange={page => this.onChangePage(page)}
        />
        }
      </Layout>
    );
  }
}

export default Home;
