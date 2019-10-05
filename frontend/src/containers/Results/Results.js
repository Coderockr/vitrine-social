import React from 'react';
import Search from '../../components/Search';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import { api } from '../../utils/api';

class Results extends React.Component {
  state = {
    loading: true,
    requests: [],
    page: 1,
    order: 'desc',
  };

  componentDidMount() {
    document.title = 'Vitrine Social - Resultados da Busca';

    this.fetchRequests();
  }

  componentDidUpdate(prevProps) {
    const { match: { params } } = this.props;
    const previousParams = prevProps.match.params;
    if (previousParams && params.searchParams !== previousParams.searchParams) {
      this.fetchRequests();
    }
  }

  onChangePage(page) {
    this.setState({
      page,
    }, () => {
      const { history } = this.props;
      const pathname = this.getPathname(history.location.pathname, 'page=', `page=${this.state.page}`);
      history.push(pathname);
    });
  }

  getPathname(pathname, param, newValue) {
    const split = pathname.split('&');
    const index = split.indexOf(split.find((value) => {
      if (value.includes(param)) {
        return true;
      }
      return false;
    }));
    split[index] = newValue;
    return split.join('&');
  }

  orderChanged(value) {
    let order = 'desc';
    if (value === 'OLDEST') {
      order = 'asc';
    }
    this.setState({
      order,
    }, () => {
      const { history } = this.props;
      const pathname = this.getPathname(history.location.pathname, 'order=', `order=${this.state.order}`);
      history.push(pathname);
    });
  }

  fetchRequests() {
    const { match: { params } } = this.props;
    const textParam = params.searchParams.split('&', 1)[0];
    let searchedtext = textParam.split('=')[1];
    if (!textParam.includes('text=')) {
      searchedtext = '';
    }
    this.setState({ text: searchedtext, loading: true });
    let search = params.searchParams;
    if (!search) {
      search = `page=${this.state.page}`;
    }
    api.get(`search?${search}`).then(
      (response) => {
        this.setState({
          requests: response.data.results,
          pagination: response.data.pagination,
          loading: false,
        });
      }, (error) => {
        this.setState({
          loading: false,
          error,
        });
      },
    );
  }

  searchRequests(text) {
    const { history } = this.props;
    history.push(`/busca/text=${text}&page=1&status=ACTIVE&orderBy=createdAt&order=${this.state.order}`);
  }

  render() {
    const { pagination } = this.state;
    return (
      <Layout>
        <Search
          text={this.state.text}
          search={text => this.searchRequests(text)}
        />
        <Requests
          loading={this.state.loading}
          activeRequests={this.state.loading ? null : this.state.requests}
          orderChanged={order => this.orderChanged(order.target.value)}
          error={this.state.error}
          search
        />
        {pagination &&
          <Pagination
            current={this.state.page}
            total={pagination.totalResults}
            onChange={page => this.onChangePage(page)}
          />
        }
      </Layout>
    );
  }
}

export default Results;
