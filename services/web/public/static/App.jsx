const {Table, Modal, Button} = ReactBootstrap;
let token = localStorage.getItem('access_token');
let apiURL;

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            adCampaigns: [],
        };
        this.loadAdCampaigns = this.loadAdCampaigns.bind(this);
    }

    // fetch ad campaigns from the web server
    loadAdCampaigns() {
        fetch(apiURL + '/api/v1/ad-campaigns', {
            method: 'GET',
            headers: new Headers({
                'Authorization': 'Bearer '+ token,
            }),
        })
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        adCampaigns: result,
                    });
                },
                (error) => {
                    // TODO handle
                }
            );
    }

    componentDidMount() {
        apiURL = document.getElementById('api').value;
        this.loadAdCampaigns();
    }

    render() {
        return <AdCampaigns data={this.state.adCampaigns}/>
    }
}

// the modal to display more details about
// the selected ad campaign
class AdCampaignModal extends React.Component {
    render() {
        const adCampaign = this.props.adCampaign;
        return (
            adCampaign && <Modal
                {...this.props}
                bsSize="large"
                aria-labelledby="contained-modal-title-lg"
                dialogClassName="ac-modal"
            >
                <Modal.Header closeButton>
                    <Modal.Title id="contained-modal-title-lg">Ad campaign with ID #{adCampaign.ID}</Modal.Title>
                </Modal.Header>
                <Modal.Body style={{}}>
                    <AdCampaign adCampaign={adCampaign} />
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={this.props.onHide}>Close</Button>
                </Modal.Footer>
            </Modal>
        );
    }
}

// table overview of the ad campaigns
class AdCampaigns extends React.Component {
    constructor(props) {
        super(props);

        this.closeModal = this.closeModal.bind(this);
        this.activateAdCampaign = this.activateAdCampaign.bind(this);

        this.state = {
            activeAdCampaign: null
        }
    }

    closeModal() {
        this.setState({
            activeAdCampaign: null
        });
    }

    activateAdCampaign(adCampaignId) {
        const filteredAdCampaigns = this.props.data.filter(ac => ac.ID === adCampaignId);
        const activeAdCampaign = filteredAdCampaigns[0] || null;
        this.setState({
            activeAdCampaign: activeAdCampaign
        });
    }

    render() {
        return (
            <React.Fragment>
                <AdCampaignModal show={!!this.state.activeAdCampaign} onHide={this.closeModal} adCampaign={this.state.activeAdCampaign}/>
                <Table responsive>
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Goal</th>
                        <th>Total budget</th>
                        <th>Status</th>
                    </tr>
                    </thead>
                    <tbody>
                    {this.props.data.map(v => {
                        return <tr>
                                <td><a onClick={() => this.activateAdCampaign(v.ID)}>#{v.ID}</a></td>
                                <td>{v.Name}</td>
                                <td>{v.Goal}</td>
                                <td>{v.TotalBudget}</td>
                                <td>{v.Status}</td>
                            </tr>
                    })}
                    </tbody>
                </Table>
            </React.Fragment>
        );
    }
}

// a detailed representation
// of a single ad campaign
class AdCampaign extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
                <div className={"row"}>
                    {Object.keys(this.props.adCampaign).map(k =>{
                        if (typeof this.props.adCampaign[k] !== 'object') {
                            return (
                                <React.Fragment>
                                    <div className={"col-md-2"}>
                                        <b>{k}</b>
                                    </div>
                                    <div className={"col-md-10"}>
                                        <div>{this.props.adCampaign[k]}</div>
                                    </div>
                                </React.Fragment>
                            )
                        }
                    })}

                    <div className={"col-md-2"}>
                        <b>Platforms</b>
                    </div>
                    <div className={"col-md-10"}>
                        <TablePlatform data={this.props.adCampaign.Platforms}/>
                    </div>
                </div>

        );
    }
}

// a way to represent a child object
// is by a table as an attribute
const TableAttr = (props) => {
    return (
        <Table responsive>
            {Object.keys(props.data).map((k) => {
                return (
                    <tr>
                        <th>{k}</th>
                        <td>{k==='Image'?<a href={'/static/image/'+props.data[k]}
                                               target="_blank"><img
                                            src={'/static/image/'+props.data[k]}
                                            height="50"
                                            width="50"/></a>:
                            k==='URL'?<a href={props.data[k]}>{props.data[k]}</a>:
                                props.data[k].toString()}</td>
                    </tr>
                )
            })}
        </Table>
    )
};

// a table with a single platform's
// attributes
const TablePlatform = (props) => {
    return (
        <Table responsive>
            <thead>
                <tr>
                    <th>
                        Attr
                    </th>
                    {Object.keys(props.data).map((k) => {
                        return (
                            <th>
                                <th>{k}</th>
                            </th>
                        )
                    })}
                </tr>
            </thead>
            <tbody>
                {Object.keys(props.data[Object.keys(props.data)[0]]).map((a) => {
                    return (
                        <tr>
                            <td>
                                <b>{a}</b>
                            </td>
                                {Object.keys(props.data).map((k) => {
                                    if (props.data[k][a] === null) {
                                        return
                                    }
                                    if (typeof props.data[k][a] === 'object') {
                                        return <td><TableAttr data={props.data[k][a]}/></td>
                                    }
                                    return (
                                        <td>{a.includes('Date')?moment(props.data[k][a]).format('DD.MM.YYYY'):props.data[k][a]}</td>
                                    )
                                })}
                        </tr>
                    )
                })}
            </tbody>

        </Table>
    )
};

ReactDOM.render(<App />, document.getElementById('root'));
