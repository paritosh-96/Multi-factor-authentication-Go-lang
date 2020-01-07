import React, {useState} from 'react';
import {withStyles} from '@material-ui/core/styles';
import Edit from "../utils/edit.png";
import Save from "../utils/saveimg.jpeg";
import Cancel from "../utils/cancelImg.png";

const useStyles = withStyles(theme => ({
    root: {
        width: '80%',
        marginTop: theme.spacing(3),
        overflowX: 'auto',
        marginLeft: '10px'
    },
}));

const TableRenderer = (props) => {
    const classes = useStyles;
    const [editingId, setEditingId] = useState("");
    const [update,setUpdate] = useState([]);

    const deleteClickHandler = (event, id) => {
        props.handleDelete(id);
    };

    const editClickHandler = (event, id) => {
        setEditingId(id);
    };

    const cancelEdit = () => {
      setEditingId("");
    };

    return (
        <div>
            <table className="question-view-table table-hover">
                <thead>
                {props.headers.map(header => <th>{header}</th>)}
                <th>Actions</th>
                </thead>
                <tbody>
                {props.data.map((row, i) =>
                    <tr>
                        {props.headers.map(h => {
                                if (h === "SerialNo" && editingId === row.QuestionId) {
                                    return (<input type="text" defaultValue={row[h]} onChange={(e) => props.handleSerialNoEdit(e, row.QuestionId, i)}/>);
                                } else {
                                    return (<td>{row[h]}</td>);
                                }
                            }
                        )}
                        <td>
                            {editingId === "" && <img className="table-edit-img" alt="edit" src={Edit} onClick={(e) => editClickHandler(e, row.QuestionId)}/>}
                            {editingId !== "" && <img className="table-edit-img" alt="save" src={Save} onClick={(e) => props.handleSave(e, row.QuestionId)}/>}
                            {editingId !== "" && <img className="table-edit-img" alt="cancel-edit" src={Cancel} onClick={cancelEdit}/>}
                            <button className="btn btn-danger" onClick={(e) => deleteClickHandler(e, row.QuestionId)}>Delete</button>
                        </td>
                    </tr>
                )}
                </tbody>
            </table>
        </div>
    );
};

export default TableRenderer;