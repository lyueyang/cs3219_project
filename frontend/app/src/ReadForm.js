import React from 'react';
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card'
import axios from 'axios';

class ReadForm extends React.Component {
    state = {
        userInfo: "",
    };

    handleRead = (event) => {
        console.log("Reading User")

        axios.get('/accounts')  
            .then((resp) => {
                const info = resp.data
                console.log(info)
                this.setState({userInfo : JSON.stringify(info)})
            })
    }

    render() {
        return (
            <div className="bg-light m-3 p-5">
                <h3 className="mb-3">Reading Users</h3>
                <Card className="mb-3">
                    <Card.Body>{this.state.userInfo}</Card.Body>
                </Card>
                
                <Button 
                    type="button"
                    onClick={this.handleRead}
                >
                    Read Users
                </Button>
            </div>
        )
    }
}

export default ReadForm;