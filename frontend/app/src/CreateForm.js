import React from 'react';
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'
import FloatingLabel from 'react-bootstrap-floating-label'
// import axios from 'axios'

class CreateForm extends React.Component {
    state = {
        username: "",
        desc: ""
    }

    handleCreate = (event) => {
        console.info("Creating User")

        var data = {
            Username: this.state.username,
            Description: this.state.desc,
        }

        fetch('/accounts', {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data),
        })
        .then(data => console.log(data))
        .catch(err => console.error(err))
    }

    render() {
        return (
            <div className="bg-light m-3 p-5">
                <h3 className="mb-3">Adding Users</h3>
                <Form>
                    <Form.Group>
                        <FloatingLabel
                            label="Username"
                            onChange={e => {
                                this.state.username = e.target.value
                            }}
                        >
                            <Form.Control type="text" placeholder="Enter Username"/>
                        </FloatingLabel>
                    </Form.Group>
                    <Form.Group>
                        <FloatingLabel
                            label="Description"
                            onChange={e => {
                                this.state.desc = e.target.value
                            }}
                        >
                            <Form.Control type="text" placeholder="Enter Description"/>
                        </FloatingLabel>
                    </Form.Group>
                    
                    <Button 
                        type="button"
                        onClick={this.handleCreate}
                    >
                        Create User
                    </Button>
                </Form>
            </div>
        )
    }
}

export default CreateForm;