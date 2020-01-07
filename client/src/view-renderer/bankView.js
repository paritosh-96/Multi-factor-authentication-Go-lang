import React, {useEffect, useState} from 'react';
import TableRenderer from '../utils/tablerenderer';
import axios from 'axios';
import {baseUrl} from "../utils/util";
import Modal from "react-bootstrap/Modal"
import Button from "react-bootstrap/Button"
import Form from "react-bootstrap/Form"
import Col from "react-bootstrap/Col"
import Row from "react-bootstrap/Row"
import swal from 'sweetalert';

const BankView = (props) => {

    const [questions, setQuestions] = useState();
    const [addingQuestion, setAddingQuestion] = useState(null);
    const [newQuestion, setNewQuestion] = useState("");
    const [update, setUpdate] = useState([]);

    let templateJson = {
        "attrFields": ["SerialNo", "Question"],
    };

    const bootstrap = () => {
        axios.get(baseUrl + 'api/bank/listAll').then(response => {
            setQuestions(response.data);
        })
    };

    const handleAddNewQuestion = () => {
        setAddingQuestion(false);
        let body = {"userId": props.userId, "question": newQuestion};
        axios.post(baseUrl + 'api/bank/addQuestion', body).then(response => {
            swal("Added successfully", "Question added successfully", "success");
            bootstrap();
        }).catch(error => {
            swal("Oops", "Error while adding question: " + (error.response.data ? error.response.data : error), "error");
        });
        setNewQuestion("");
    };

    const handleQuestionDelete = (questionId) => {
        let body = {"questionId": Number(questionId), "userId": props.userId};
        axios.post(baseUrl + 'api/bank/disableQuestion', body).then(response => {
            swal("Deleted Successfully", "", "success");
            bootstrap()
        }).catch(error => {
            swal("Oops", "Error while deleting question: " + error, "error");
        })
    };

    const updateSerialNo = (event, id, ind) => {
        let _update = [];
        _update[id] = event.target.value;
        setUpdate(_update);
    };

    const saveClickHandler = (event, id) => {
        console.log("Update serial no of :", id, "as : ", update[id]);

    };

    const handleNewQuestion = (event) => setNewQuestion(event.target.value);
    const handleClose = () => setAddingQuestion(false);
    const handleShow = () => setAddingQuestion(true);

    useEffect(() => {
        bootstrap();
    }, []);

    return (
        <div>
            <p className="welcome-message">Welcome <span className="user-name">{props.userId}</span>, Work on the security questions of the Bank</p>
            <h3>Questions:</h3>
            <div className="Bank-question-list">
                <div className="question-view-div">
                    {questions && <TableRenderer data={questions} handleDelete={handleQuestionDelete}
                                                 handleSerialNoEdit={updateSerialNo} handleSave={saveClickHandler}
                                                 headers={templateJson.attrFields}/>}
                </div>
                <button type="button" className="btn btn-info btn-lg add-btn" onClick={handleShow}>
                    Add Question
                </button>
                {addingQuestion === true &&
                <Modal show={addingQuestion} onHide={handleClose} centered>
                    <Modal.Header closeButton>
                        <Modal.Title>Add a new question</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Form>
                            <Form.Group as={Row} controlId="formPlaintextEmail">
                                <Form.Label column sm="2">Question:</Form.Label>
                                <Col sm="10">
                                    <Form.Control placeHolder="Enter the question" onChange={handleNewQuestion}/>
                                </Col>
                            </Form.Group>
                        </Form>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={handleClose}>
                            Close
                        </Button>
                        <Button variant="primary" onClick={handleAddNewQuestion}>
                            Save Changes
                        </Button>
                    </Modal.Footer>
                </Modal>
                }
            </div>
        </div>
    );
};

export default BankView