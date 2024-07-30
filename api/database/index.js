const users = require('./users.json');
const transactions = require('./transactions.json');

module.exports = function () {
    return {
        users: users,
        transactions: transactions,
    }
}
