import React from 'react';
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'
import FloatingLabel from 'react-bootstrap-floating-label'

class DeleteForm extends React.Component { 
    state = {
        user: ""
    }

    handleDelete = (event) => {
        event.preventDefault();
        console.log("Deleting User")
        
        var data = {
            Username: this.state.user
        }

        fetch('/accounts', {
            method: 'DELETE',
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
                <h3 className="mb-3">Delete Users</h3>
                <Form>
                    <Form.Group>
                        <FloatingLabel
                            label="Username"
                            onChange={e => {
                                this.state.user = e.target.value
                            }}
                        >
                            <Form.Control 
                                className="textFeedback"
                                type="text" 
                                name="username" 
                                required
                                />
                        </FloatingLabel>
                    </Form.Group>
                    
                    <Button 
                        type="button"
                        onClick= {this.handleDelete} 
                    >
                        Delete User
                    </Button>
                </Form>
            </div>
        )
    }
}

export default DeleteForm;