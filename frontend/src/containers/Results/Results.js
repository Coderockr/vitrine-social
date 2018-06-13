import React from 'react';
import Search from '../../components/Search';
import Pagination from '../../components/Pagination';
import Layout from '../../components/Layout';
import Requests from '../../components/Requests';
import api from '../../utils/api';

class Results extends React.Component {
  state = {
    loading: true,
    requests: [],
    page: 1,
  }

  componentWillMount() {
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
      history.push(`/search/text=${this.state.text}&page=${this.state.page}&status=ACTIVE`);
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
      },
    );
  }

  searchRequests(text) {
    const { history } = this.props;
    history.push(`/search/text=${text}&page=1&status=ACTIVE`);
  }

  render() {
    return (
      <Layout>
        <Search
          text={this.state.text}
          search={text => this.searchRequests(text)}
        />
        <Requests
          loading={this.state.loading}
          activeRequests={this.state.loading ? null : this.state.requests}
          search
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

export default Results;
