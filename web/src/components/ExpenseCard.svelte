<script lang="ts">
  import type {Expense} from "../services/expenses";
  import { deleteExpense } from "../services/expenses";
  import { 
    Card, 
    CardText, 
    CardActions, 
    Button, 
    MaterialApp,
    Icon
  } from 'svelte-materialify';
  import { createEventDispatcher } from "svelte";
  import { mdiTrashCan, mdiPen } from "@mdi/js";

  export let expense: Expense;
  const dispatch = createEventDispatcher();

  const handleDelete = () => {
    const res = deleteExpense(expense.id);
    res.then(() => dispatch('remove', expense));
  }

  const dispatchUpdate = () => {
    dispatch("update", expense);
  }
</script>

<MaterialApp>
  <div class="d-flex justify-center mt-4 mb-4">
    <Card raised style="width: 700px;">
      <div class="pl-4 pr-4 pt-3 d-flex justify-space-between">
        <div>
          <span class="text-overline">{expense.category}</span>
          <br />
          <span class="text-h5 mb-2">
            {expense.amount}
          </span>
          <br />
          <CardText style="padding: 16px 16px 16px 0px;">
            {expense.description}
          </CardText>
        </div>
        <div>
          <CardActions>
            <Button class="mr-3 primary-color" on:click={handleDelete}>
              <Icon path={mdiTrashCan} />
            </Button>
            <Button class="primary-color" on:click={dispatchUpdate}>
              <Icon path={mdiPen} />
            </Button>
          </CardActions>
        </div>
      </div>
    </Card>
  </div>
</MaterialApp>