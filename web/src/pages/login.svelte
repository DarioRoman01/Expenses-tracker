<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import { redirect, url } from "@roxi/routify";
  import { Row, Col, TextField, Button, MaterialApp } from 'svelte-materialify';

  let usernameOrEmail: string;
  let password: string;
  let error: Error;

  const usernameRules = [(v: any) => !!v || 'Required'];
  const passwordRules = [
    (v: any) => !!v || 'Required',
    (v: string|any[]) => v.length <= 4 || 'must be 4 characters',
  ]

  async function login() {
    try {
      const res = await fetch("http://localhost:1323/login", {
        method: "post",
        body: JSON.stringify({
          usernameOrEmail: usernameOrEmail,
          password: password,
        }),
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (res.ok) {
        $redirect("./home");
      } else {
        const err = await res.json();
        console.log(err)
        error = new Error(err.message);
      }
    } catch (err) {
      console.log(err)
      error = new Error(err);
    }
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
        <Button size="large" class="primary-color" on:click={login}>
          Submit
        </Button>
      </div>
    </div>
  </div>
</MaterialApp>
