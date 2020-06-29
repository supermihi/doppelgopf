import React from 'react';
import { TableState } from 'model/table';
import AddIcon from '@material-ui/icons/Add';
import SearchIcon from '@material-ui/icons/Search';
import MailOutlineIcon from '@material-ui/icons/MailOutline';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import makeStyles from '@material-ui/core/styles/makeStyles';
import CurrentTableView from './CurrentTableView';

interface Props {
  activeTable: TableState | null;
  createTable: () => void;
}

const useStyles = makeStyles((theme) => ({
  addTable: {
    marginTop: theme.spacing(2),
    alignSelf: 'center',
  } as const,
  main: {
    display: 'flex',
    marginTop: theme.spacing(2),
    flexDirection: 'column',
  } as const,
  buttons: {
    marginTop: theme.spacing(2),
  },
}));

export default ({ activeTable, createTable }: Props) => {
  const classes = useStyles();
  return (
    <div className={classes.main}>
      {activeTable && <CurrentTableView table={activeTable} />}
      <Grid container spacing={2} className={classes.buttons}>
        <Grid item xs={4}>
          <Button fullWidth startIcon={<MailOutlineIcon />}>
            Einladung
          </Button>
        </Grid>
        <Grid item xs={4}>
          <Button fullWidth startIcon={<SearchIcon />}>
            Tisch suchen
          </Button>
        </Grid>
        <Grid item xs={4}>
          <Button startIcon={<AddIcon />} fullWidth onClick={createTable}>
            Neuer Tisch
          </Button>
        </Grid>
      </Grid>
    </div>
  );
};
