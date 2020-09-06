<template>
  <v-simple-table>
    <thead>
      <tr>
        <th class="text-left">URL</th>
        <th class="text-left">Request Type</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><slot></slot></td>
        <td>{{type}}</td>
      </tr>
    </tbody>
  </v-simple-table>
</template>

<script>
export default {
  name: "apitable",
  props: ["type"],
}
</script>
