import React, {useEffect, useState} from 'react';
import axios from 'axios';
import {baseUrl} from "../utils/util";
import Form from "react-bootstrap/Form";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";
import swal from "sweetalert";

const ValidationView = (props) => {
    const [challengeQuestions, setChallengeQuestions] = useState([]);
    const [questionAnswer, setQuestionAnswer] = useState([]);

    const bootstrap = () => {
        axios.get(baseUrl + 'api/event/challenge?customerId=' + props.userId).then(response => {
            setChallengeQuestions(response.data);
            let _qA = [];
            for (let i = 0; i < response.data.length; i++) {
                _qA.push({...response.data[i], CustomerId: props.userId});
            }
            setQuestionAnswer(_qA);
            console.log(" Q A:", _qA);
        }, error => {
            console.log("Error while fetching questions for customer" + props.userId + ": " + (error.response.data ? error.response.data : error));
            setQuestionAnswer([]);
            setChallengeQuestions([]);
        });
    };

    const handleAnswerChange = (id, ind, event) => {
        let _qA = JSON.parse(JSON.stringify(questionAnswer));
        _qA[ind]["QuestionId"] = parseInt(id);
        _qA[ind]["Answer"] = event.target.value;
        setQuestionAnswer(_qA);
    };

    const isAnswerValid = () => {
        for (let ques of questionAnswer) {
            if (ques["Answer"].length < 3) {
                swal("Oops", "Answers should be minimum 3 characters", "error");
                return false;
            }
        }
        return true;
    };

    const ValidateAnswers = (event) => {
        event.preventDefault();
        if (isAnswerValid()) {
            let json = JSON.stringify(questionAnswer);
            axios.post(baseUrl + 'api/event/validate', json).then(response => {
                swal("Success", "Validated successfully", "success");
                bootstrap();
            }, error => {
                swal("Error", "Invalid answers, Try again: " + (error.response.data ? error.response.data : error), "error");
            })
        }
    };

    useEffect(() => {
        bootstrap();
    }, [props.userId, props.viewChanged]);

    return (
        <div>
            <p className="welcome-message"> Welcome <span className="user-name"> {props.userId} </span>, Validate the customer answers </p>
            {challengeQuestions.length > 0 &&
            <Form onSubmit={ValidateAnswers}>
                {challengeQuestions.map((ques, i) =>
                    <Form.Row>
                        <Form.Label column sm="5"> {ques.Question} </Form.Label>
                        <Col sm="5">
                            <Form.Control id={'ans' + i} onChange={(e) => handleAnswerChange(ques.QuestionId, i, e)}
                                          placeholder="Answer..."/>
                        </Col> </Form.Row>)}
                <Form.Row>
                    <Col sm={{span: 10, offset: 1}}>
                        <Button type="submit" className="validate-button"> Validate </Button>
                    </Col>
                </Form.Row>
            </Form>}
        </div>
    );
};

export default ValidationView;