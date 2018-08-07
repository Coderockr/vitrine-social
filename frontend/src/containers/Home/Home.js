import React from 'react';
import Search from '../../components/Search';
import Categories from '../../components/Categories';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import RequestDetails from '../../components/RequestDetails';
import { api } from '../../utils/api';

class Home extends React.Component {
  state = {
    loadingCategories: true,
    loadingRequests: true,
    categories: [],
    requests: [],
    page: 1,
    order: 'desc',
  }

  componentWillMount() {
    this.fetchCategories();
    this.fetchRequests();
  }

  componentDidMount() {
    document.title = 'Vitrine Social';
    const { match: { params } } = this.props;
    if (params.requestId) {
      api.get(`need/${params.requestId}`).then(
        (response) => {
          this.setState({
            request: response.data,
          });
        },
      );
    }
  }

  onCancel() {
    const { history } = this.props;
    history.push('/');
    this.setState({ request: null });
  }

  onChangePage(page) {
    this.setState({
      page,
      loadingRequests: true,
    }, () => {
      this.fetchRequests();
    });
  }

  orderChanged(value) {
    let order = 'desc';
    if (value === 'OLDEST') {
      order = 'asc';
    }
    this.setState({ order }, () => this.fetchRequests());
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
    const { page, order } = this.state;
    api.get(`search?page=${page}&status=ACTIVE&orderBy=createdAt&order=${order}`).then(
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
    history.push(`/busca/text=${text}&page=1&status=ACTIVE&orderBy=createdAt&order=desc`);
  }

  searchByCategory(categoryId) {
    const { history } = this.props;
    history.push(`/busca/categories=${categoryId}&page=1&status=ACTIVE&orderBy=createdAt&order=desc`);
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
          orderChanged={order => this.orderChanged(order.target.value)}
        />
        {this.state.pagination &&
          <Pagination
            current={this.state.page}
            total={this.state.pagination.totalResults}
            onChange={page => this.onChangePage(page)}
          />
        }
        {this.state.request &&
          <RequestDetails
            visible={this.state.request}
            onCancel={() => this.onCancel()}
            request={this.state.request}
          />
        }
      </Layout>
    );
  }
}

export default Home;
