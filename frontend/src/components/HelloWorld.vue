<template>
  <v-container class="fill-height" fluid>
    <v-row justify="center" align="center" class="fill-height">
      <v-col cols="12" sm="10" md="8" lg="6">
        <v-card class="text-center pa-5" elevation="2" color="darkgrey" rounded="lg" outlined>
          <div>
            <h1 class="display-1 font-weight-bold">🐙.☕️</h1>
            <h2 class="headline font-weight-light">Password Checker</h2>
            <v-btn icon :href="'https://github.com/nilgaar/PassCheck'" target="_blank">
              <v-icon>mdi-github</v-icon>
            </v-btn>
          </div>

          <v-divider class="my-4"></v-divider>

          <v-card outlined color="surface-light" class="mb-4">
            <v-card-text>
              Check if your password is present in any popular password dictionary.
              By analyzing hashed versions of passwords against common dictionaries and wordlists, it helps
              determine if your password is still secure.
            </v-card-text>
          </v-card>

          <v-form @submit.prevent="hashPassword">
            <v-text-field v-model="password" label="Password" outlined type="password" required></v-text-field>
            <v-btn color="primary" class="mt-3" type="submit">Check Password</v-btn>
          </v-form>

          <v-overlay opacity=".12" scrim="primary" :value="true"></v-overlay>
        </v-card>
      </v-col>

      <v-col cols="12" sm="10" md="8" lg="6">
        <v-card class="pa-4" color="darkgrey" elevation="2" rounded="lg" outlined>
          <v-list dense>
            <v-subheader>Dictionaries containing the Password:</v-subheader>
            <v-list-item-group v-for="file in files" :key="file">
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>{{ file }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>


<script setup lang="ts">
import { ref } from "vue";
import { sha3_512 } from "js-sha3";

const password = ref("");
const files = ref([]);

async function hashPassword() {
  const hashedPassword = sha3_512(password.value);
  console.log(hashedPassword);
  fetch("https://api.checker.pops.cafe?hash=" + hashedPassword)
    .then((response) => response.json())
    .then((data) => {
      files.value = data.data.split(', ').sort().map((file: string) => file.trim());
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}
</script>

<style scoped>
/* Base text styling for white appearance */
.text-caption,
.headline,
p,
h1,
h2,
.display-1 {
  color: #FFFFFF;
}

.v-btn {
  margin-top: 12px;
}

/* Responsive settings for better readability */
@media (max-width: 600px) {

  .headline,
  .display-1,
  .text-caption {
    font-size: smaller;
  }
}
</style>
