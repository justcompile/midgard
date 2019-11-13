'use strict';

const e = React.createElement;

class WorkerSection extends React.Component {
    ws = new WebSocket("ws://localhost:8000/ws");

    constructor(props) {
        super(props);
        this.state = {workers: []}
    }

    componentDidMount() {
        this.ws.onopen = () => {
        // on connecting, do nothing but log it to the console
            console.log('connected')
        };

        this.ws.onmessage = evt => {
            // listen to data sent from the websocket server
            const message = JSON.parse(evt.data);

            console.log(message)
            
            this.upsertWorkers(message);
            
        }

        this.ws.onclose = () => {
            console.log('disconnected')
            // automatically try to reconnect on connection loss
        }
    }

    upsertWorkers(data) {
        let {workers} = this.state

        if (!Array.isArray(data)) {
            data = [data]
        }

        for (var i = 0, l = data.length; i < l; i++) {
            var idx = -1;

            for (var j = 0; j < workers.length; j++) {
                if ( data[i].worker.name == workers[j].worker.name) {
                    idx = j;
                    break;
                }
            }

            if (idx === -1) {
                workers.push(data[i])
            } else {
                workers[idx] = data[i];
            }
        }

        this.setState({workers: workers});
    }

    render() {

        return (
            <div className="sidebar-sticky">
                <h6
                    className="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted">
                    <span>Workers</span>
                    <a className="d-flex align-items-center text-muted" href="#">
                        <span data-feather="plus-circle"></span>
                    </a>
                </h6>
                <WorkerList workers={this.state.workers} />
            </div>
        );
    }
}


class WorkerList extends React.Component {
    createWorkerItems() {
        const {workers} = this.props;
        
        let items = []

        for (let i = 0, l = workers.length; i < l; i++) {
            items.push(<WorkerItem item={workers[i]} key={`worker-${i}`}/>)
        }

        return items
    }

    render() {
        return (
            <ul id="workers" className="nav flex-column mb-2">
                {this.createWorkerItems()}
            </ul>
        );
    }
}

class WorkerItem extends React.Component {
    render() {
        const {item} = this.props;

        return (
            <li className="nav-item" key={item.worker.name}>
                <a className="nav-link">
                    <span>{item.worker.name}</span>
                    <span>[{item.type}]</span>
                </a>
            </li>
        )
    }
}

const workerContainer = document.querySelector('.sidebar');
ReactDOM.render(e(WorkerSection), workerContainer);
