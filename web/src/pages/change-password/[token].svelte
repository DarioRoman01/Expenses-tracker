<script lang="ts">
  import Navbar from "../../components/Navbar2.svelte";
  import { changePassword } from "../../services/users";
  import {params, redirect} from "@roxi/routify";
  import {TextField, Button, MaterialApp, Alert } from 'svelte-materialify';

  let newPassword: string;
  let error: Error;
  
  const handleClick = () => {
    const user = changePassword($params["token"], newPassword);
    user.then(() => $redirect("../home")).catch((err: Error) => error = err)
  }
</script>

<MaterialApp>
  <Navbar isLoggedIn={false}/>
  <div class="d-flex justify-center mt-4">
    <div class="d-flex flex-column">
      <h3 class="text-h3 mb-6">Login</h3>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField bind:value={newPassword}>
          new password
        </TextField>
      </div>
      <div style="align-self: center;" class="mt-4">
        <Button size="large" class="primary-color" on:click={handleClick}>
          Submit
        </Button>
      </div>
      {#if error}
        <Alert class="error-color">
          {error.message}
        </Alert>
      {/if}
    </div>
  </div>
</MaterialApp>
