import React, {useEffect, useState} from 'react';
import Form from "react-bootstrap/Form";
import Col from "react-bootstrap/Col";
import axios from "axios";
import {baseUrl} from "../utils/util";
import Button from "react-bootstrap/Button";
import swal from "sweetalert";

const CustomerView = (props) => {
    const [allQuestion, setAllQuestion] = useState([]);
    const [noOfQuestionsPerCust, setNoOfQuestionsPerCust] = useState(0);
    const [questionAnswer, setQuestionAnswer] = useState([]);
    const [newUser, setNewUser] = useState(true);
    const [typeOfSubmit, setTypeOfSubmit] = useState("");

    const bootstrap = () => {
        axios.get(baseUrl + 'api/customer/customerQuestionCount').then(response => {
            setNoOfQuestionsPerCust(response.data);
            let _num = response.data;
            axios.get(baseUrl + 'api/customer/list?customerId=' + props.userId).then(response => {
                let _qA = [];
                if (response.data.length > 0) {
                    setNewUser(false);
                    _qA = JSON.parse(JSON.stringify(response.data));
                } else {
                    setNewUser(true);
                    for (let ind = 0; ind < _num; ind++) {
                        _qA.push({QuestionId: -1, CustomerId: props.userId, Answer: ""});
                    }
                }
                setQuestionAnswer(_qA);
            });
        });
        axios.get(baseUrl + 'api/customer/listAll').then(response => {
            let _ques = [];
            for (let ind = 0; ind < response.data.length; ind++) {
                _ques.push({...response.data[ind], disabled: false});
            }
            setAllQuestion(_ques);
        });
    };

    const handleSelectedQuestion = () => {
        let _q = JSON.parse(JSON.stringify(allQuestion));
        for (let i = 0; i < _q.length; i++) {
            let found = false;
            for (let j = 0; j < questionAnswer.length; j++) {
                if (questionAnswer[j]["QuestionId"] === _q[i].QuestionId) {
                    _q[i]["disabled"] = true;
                    found = true;
                    break;
                }
            }
            if (!found) {
                _q[i]["disabled"] = false;
            }
        }
        setAllQuestion(_q)
    };

    const updateAnswerField = (ind, event) => {
        console.log("Chnaging answers");
        let _q = JSON.parse(JSON.stringify(questionAnswer));
        _q[ind]["Answer"] = event.target.value;
        setQuestionAnswer(_q);
    };
    const updateQuestionField = (ind, event) => {
        let _q = questionAnswer;
        _q[ind]["QuestionId"] = parseInt(event.target.value);
        setQuestionAnswer(_q);
        handleSelectedQuestion();
    };

    const isValidData = () => {
        let notValid = false;
        for (let ques of questionAnswer) {
            if (ques.QuestionId === -1 || ques.Answer.length < 3) {
                notValid = true;
                break;
            }
        }
        if (notValid) {
            swal("Invalid", "The questions or answers are invalid. All questions must be answered, and the answer should be of minimum 3 characters", "error");
            return false;
        }
        return true;
    };

    const handleFormSubmit = (event) => {
        event.preventDefault();
        if (typeOfSubmit === "reset") {
            axios.post(baseUrl + 'api/customer/reset?customerId=' + props.userId).then(response => {
                swal("Success", "All answers has been reset", "success");
                bootstrap();
            }, error => {
                swal("Error", "Could not reset the answers: " + error, "error");
            })
        } else {
            if (isValidData()) {
                let json = JSON.stringify(questionAnswer);

                if (typeOfSubmit === "add") {
                    axios.post(baseUrl + 'api/customer/add', json).then(response => {
                        swal("Added successfully", "All the answers were added successfully", "success");
                        bootstrap();
                    }, error => {
                        swal("Error", "Error while adding answers , " + error + " Please check logs", "error")
                    });
                } else {
                    axios.post(baseUrl + 'api/customer/modify', json).then(response => {
                        swal("Modified successfully", "Modification Status:" + JSON.stringify(response.data), "success");
                        bootstrap();
                    }, error => {
                        swal("Error", "Error while modifying answers, " + error + " Please check logs", "error")
                    })
                }
            }
        }
    };

    const handleAddSubmit = () => {
        setTypeOfSubmit("add");
    };

    const handleModifySubmit = () => {
        setTypeOfSubmit("modify");
    };

    const handleResetSubmit = () => {
        setTypeOfSubmit("reset");
    };

    useEffect(() => {
        bootstrap();
    }, [props.userId]);

    return (
        <div>
            <p>Welcome <span className="user-name"> {props.userId}</span>, Add security questions for the customer</p>
            <div className="CustomerQuestionView">
                {noOfQuestionsPerCust !== 0 &&
                <Form id="questionForm" onSubmit={handleFormSubmit}>
                    {questionAnswer.map((quesAns, i) =>
                        <Form.Row>
                            <Col>
                                <Form.Control as="select" id={'ques' + i} onChange={(e) => updateQuestionField(i, e)}
                                              disabled={!newUser}
                                              value={quesAns["QuestionId"]}>
                                    <option disabled value={-1}>Choose a question...</option>
                                    {allQuestion.map(ques => {
                                        if (ques.disabled) return (
                                            <option disabled={true} value={ques.QuestionId}>{ques.Question}</option>);
                                        else return (<option value={ques.QuestionId}>{ques.Question}</option>);

                                    })}
                                </Form.Control>
                            </Col>
                            <Col>
                                <Form.Control id={'ans' + i} onChange={(e) => updateAnswerField(i, e)}
                                              value={quesAns["Answer"]}
                                              placeholder="Answer..."/>
                            </Col>
                        </Form.Row>
                    )}
                    <div className="btn-div">
                        {newUser && <button className="btn btn-primary btn-add" type="submit" onClick={handleAddSubmit}>Add Questions</button>}
                        {!newUser && <button className="btn btn-primary btn-modify" type="submit" onClick={handleModifySubmit}>Modify Questions</button>}
                        {!newUser && <button className="btn btn-primary btn-reset" type="submit" onClick={handleResetSubmit}>Reset</button>}
                    </div>
                </Form>}

            </div>
        </div>)

};

export default CustomerView;