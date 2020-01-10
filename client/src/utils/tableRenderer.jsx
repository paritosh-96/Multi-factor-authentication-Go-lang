import React, {useEffect, useState} from 'react';
import {withStyles} from '@material-ui/core/styles';
import {MdDelete} from 'react-icons/md';
import $ from "jquery";
import swal from "sweetalert"

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
    const [update, setUpdate] = useState([]);

    const deleteClickHandler = (event, id) => {
        swal({
            title: "Delete",
            text: "Are you sure you want to delete this questions?",
            icon: "warning",
            buttons: true,
            dangerMode: true
        }).then((willDelete) => {
            if(willDelete) {
                props.handleDelete(id);
            }
        });
    };

    const editClickHandler = (event, id) => {
        setEditingId(id);
    };

    const cancelEdit = () => {
        setEditingId("");
    };

    useEffect(() => {
        $('#question-view-tbl > tbody').sortable({
            stop: function (event, ui) {
                debugger
                console.log(event, ui.position);
            }
        }).disableSelection();
    });

    return (
        <div>
            <table id="question-view-tbl" className="question-view-table table-hover">
                <thead>
                {props.headers.map(header => <th> {header}</th>)}
                    <th> Actions</th>
                </thead>
                <tbody>
                {props.data.map((row, i) =>
                    <tr>
                        {props.headers.map(h => <td> {row[h]}</td>)}
                        <td>
                            <span className="delete-button" title="Delete"
                                  onClick={(e) => deleteClickHandler(e, row.QuestionId)}>
                                <MdDelete/>
                            </span>
                        </td>
                    </tr>)}
                </tbody>
            </table>
        </div>
    );
};

export default TableRenderer;

