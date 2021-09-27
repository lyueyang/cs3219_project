import React from 'react';
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'
import FloatingLabel from 'react-bootstrap-floating-label'

class UpdateForm extends React.Component {
    state = {
        username: "",
        desc: ""
    }

    handleUpdate = (event) => {
        event.preventDefault();
        console.log("Deleting User")
        var data = {
            Username: this.state.username,
            Description: this.state.desc,
        }

        fetch('/accounts', {
            method: 'PUT',
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
                <h3 className="mb-3">Update Users</h3>
                <Form>
                    <Form.Group>
                        <FloatingLabel
                            label="Username"
                            onChange={e => {
                                this.state.username = e.target.value
                            }}
                        >
                            <Form.Control type="text" />
                        </FloatingLabel>
                    </Form.Group>
                    <Form.Group>
                        <FloatingLabel
                            label="Description"
                            onChange={e => {
                                this.state.desc = e.target.value
                            }}
                        >
                            <Form.Control type="text" />
                        </FloatingLabel>
                    </Form.Group>
                    
                    <Button 
                        type="button"
                        onClick={this.handleUpdate}
                    >
                        Submit
                    </Button>
                </Form>
            </div>
        )
    }
}

export default UpdateForm;