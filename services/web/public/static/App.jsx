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
                    <Modal.Title id="contained-modal-title-lg">Ad campaign with ID #{adCampaign.id}</Modal.Title>
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
        const filteredAdCampaigns = this.props.data.filter(ac => ac.id == adCampaignId);
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
                                <td><a onClick={() => this.activateAdCampaign(v.id)}>#{v.id}</a></td>
                                <td>{v.name}</td>
                                <td>{v.goal}</td>
                                <td>{v.total_budget}</td>
                                <td>{v.status}</td>
                            </tr>
                    })}
                    </tbody>
                </Table>
            </React.Fragment>
        );
    }
}

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
                                        <b style={{"text-transform": "capitalize"}}>{k.split('_').join(' ')}</b>
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
                        <TablePlatform data={this.props.adCampaign.platforms}/>
                    </div>
                </div>

        );
    }
}

const TableAttr = (props) => {
    return (
        <Table responsive>
            {Object.keys(props.data).map((k) => {
                return (
                    <tr>
                        <th style={{"text-transform": "capitalize"}}>{k.split('_').join(' ')}</th>
                        <td>{k==='image'?<img src={props.data[k]}/>:props.data[k].toString()}</td>
                    </tr>
                )
            })}
        </Table>
    )
};

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
                                <th style={{"text-transform": "capitalize"}}>{k}</th>
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
                                <b style={{"text-transform": "capitalize"}}>{a.split('_').join(' ')}</b>
                            </td>
                                {Object.keys(props.data).map((k) => {
                                    if (!props.data[k][a]) {
                                        return
                                    }
                                    if (typeof props.data[k][a] === 'object') {
                                        return <td><TableAttr data={props.data[k][a]}/></td>
                                    }
                                    return (
                                        <td>{props.data[k][a]}</td>
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
//registerServiceWorker();