'use strict';

const e = React.createElement;
const Button = ReactBootstrap.Button;
const Modal = ReactBootstrap.Modal;

class ProjectList extends React.Component {
    constructor(props) {
        super(props);
    }

    renderProjects() {
        let items = [];
        const {projects} = this.props

        for(let i = 0, l = projects.length; i < l; i++) {
            items.push((
                <li key={`project-${i}`} className="list-group-item">{projects[i].name}</li>
            ))
        }

        return items;
    }

    render() {
        return (
            <div className="row">
                <div className="card px-0 col-12">
                    <ul className="list-group list-group-flush">{this.renderProjects()}</ul>
                </div>
            </div>
        )
    }
}

class NewProject extends React.Component {
    constructor(props) {
        super(props)
       // this.handleClose = this.handleClose.bind(this);
        this.submit = this.submit.bind(this);
    }

    handleClose() {
        const {onClose} = this.props

        onClose();
    }

    submit() {
        console.log("write me ")
    }

    render() {   
        const {show, onClose} = this.props
        return (
            <Modal show={show} onHide={onClose}>
                <Modal.Header closeButton>
                <Modal.Title>Modal heading</Modal.Title>
                </Modal.Header>
                <Modal.Body>Woohoo, you're reading this text in a modal!</Modal.Body>
                <Modal.Footer>
                <Button variant="secondary" onClick={onClose}>
                    Close
                </Button>
                <Button variant="primary" onClick={this.submit}>
                    Save Changes
                </Button>
                </Modal.Footer>
            </Modal>
        )
    }
}

class Main extends React.Component {
    constructor(props) {
        super(props)
        this.state = {projects: [],modalOpen: false}

        this.openProjectModal = this.openProjectModal.bind(this);
        this.fetchProjects = this.fetchProjects.bind(this);
        this.closeModal = this.closeModal.bind(this);
    }

    componentDidMount() {
        this.fetchProjects()
    }

    fetchProjects() {
        let that = this;

        fetch("/api/projects")
        .then(function(response) {
            if (!response.ok) {
                throw new Error(response.statusText);
            }

            return response.json()
        })
        .then(function(data) {
            that.setState({projects: data});
        });
    }

    openProjectModal() {
        this.setState({modalOpen: true})
    }

    closeModal() {
        this.fetchProjects()
        this.setState({modalOpen: false})
    }

    render() {
        return (
            <React.Fragment>
                <NewProject show={this.state.modalOpen} onClose={this.closeModal} />
                <div className="row justify-content-end mt-2">
                    <button type="button" className="btn btn-success" onClick={this.openProjectModal}>New Project</button>
                </div>
                <div className="row">
                    <h4 className="pb-2">Projects</h4>
                </div>
                <ProjectList projects={this.state.projects}/>
            </React.Fragment>
        )
    }
}


const mainContainer = document.querySelector('#main-panel');
ReactDOM.render(e(Main), mainContainer);
