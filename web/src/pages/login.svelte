<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import { login } from "../services/users";
  import { redirect, url } from "@roxi/routify";
  import {TextField, Button, MaterialApp, Alert } from 'svelte-materialify';

  let usernameOrEmail: string;
  let password: string;
  let error: Error;

  const usernameRules = [(v: any) => !!v || 'Required'];
  const passwordRules = [
    (v: any) => !!v || 'Required',
    (v: string|any[]) => v.length <= 4 || 'must be 4 characters',
  ]

  const handleLogin = () => {
    const user = login({usernameOrEmail, password});
    user.then(() => $redirect("./home")).catch((err: Error) => error = err);
  }
</script>

<MaterialApp>
  <Navbar isLoggedIn={false}/>
  <div class="d-flex justify-center mt-4">
    <div class="d-flex flex-column">
      <h3 class="text-h3 mb-6">Login</h3>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField bind:value={usernameOrEmail} rules={usernameRules}>
          username
        </TextField>
      </div>
      <div style="width: 700px;">
        <TextField type="password" bind:value={password} rules={passwordRules}>
          password
        </TextField>
      </div>
      <div style="align-self: center;" class="mt-4">
        <Button size="large" class="primary-color" on:click={handleLogin}>
          Submit
        </Button>
        <a href={$url("./forgot-password")} class="text-decoration-none ml-3">
          forgot password?
        </a>
      </div>
      {#if error}
        <Alert class="error-color">
          {error.message}
        </Alert>
      {/if}
    </div>
  </div>
</MaterialApp>
