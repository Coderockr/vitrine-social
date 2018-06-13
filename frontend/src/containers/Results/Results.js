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

  fetchRequests() {
    const { match: { params } } = this.props;
    let search = params.searchParams;
    if (!search) {
      search = `page=${this.state.page}`;
    }
    api.get(`search?${search}`).then(
      (response) => {
        this.setState({
          requests: response.data.results,
          loading: false,
        });
      },
    );
  }

  searchRequests(text) {
    const { history } = this.props;
    history.push(`/search/text=${text}&page=1`);
  }

  render() {
    const { match: { params } } = this.props;
    const searchedtext = params.searchParams.split('&', 1)[0].split('=')[1];
    return (
      <Layout>
        <Search
          text={searchedtext}
          search={text => this.searchRequests(text)}
        />
        <Requests
          loading={this.state.loading}
          activeRequests={this.state.loading ? null : this.state.requests}
          search
        />
        {!this.state.loading && <Pagination />}
      </Layout>
    );
  }
}

export default Results;
