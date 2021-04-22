<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import {TextField, Button, MaterialApp, Alert } from 'svelte-materialify';
  import { redirect, url } from "@roxi/routify";
  import { register, emailRegex, CustomError } from "../services/users";

  let username: string;
  let email: string;
  let password: string;
  let error: CustomError;

  const usernameRules = [
    (v: string) => !!v || 'Required',
    (v: string) => v.length >= 3 || 'must be at least 3 characters',
  ];

  const emailRules = [
    (v: string) => !!v || 'Required',
    (v: string) => {
      const pattern = emailRegex;
      return pattern.test(v) || 'Invalid e-mail.';
    },
  ]

  const passwordRules = [
    (v: string) => !!v || 'Required',
    (v: string) => v.length >= 4 || 'must be at least 4 characters',
  
  ]

  const handleRegister = () => {
    const user = register({username, email, password});
    user.then(() => $redirect("./home")).catch((err: CustomError) => error = err)
  }

</script>

<MaterialApp>
  <Navbar isLoggedIn={false}/>
  <div class="d-flex justify-center mt-4">
    <div class="d-flex flex-column">
      <h3 class="text-h3 mb-6">Register</h3>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField 
          error={error && error.field === "username" ? true : false} 
          bind:value={username} 
          rules={usernameRules}>
          username
        </TextField>
      </div>
      <div style="width: 700px;" class="mb-4 mt-3">
        <TextField 
          error={error && error.field === "email" ? true : false}
          bind:value={email} 
          rules={emailRules}
        >
          email
        </TextField>
      </div>
      <div style="width: 700px;">
        <TextField 
          error={error && error.field === "password" ? true : false} 
          type="password" 
          rules={passwordRules} 
          bind:value={password}
        >
          password
        </TextField>
      </div>
      <div style="align-self: center;" class="mt-4">
        <Button size="large" class="primary-color" on:click={handleRegister}>
          Submit
        </Button>
      </div>
      {#if error}
        <Alert class="error-color mt-4">
          {error.message}
        </Alert>
      {/if}
    </div>
  </div>
</MaterialApp>
