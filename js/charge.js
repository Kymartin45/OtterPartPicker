var stripe = Stripe('pk_test_51Hpq4CIkuecL7QkUYJy3ab2YCNTeB63EbUF2SiUw2II12HY4DbA391Vf3t9pFxLjsrxJayaWkCLVPybVk4fi4ume00OTiZijKd')
var elements = stripe.elements();

var style = {
  base: {
    fontSize: '16px',
    color: '#32325d',
  },
};

// Create a token or display an error when the form is submitted.
var form = document.getElementById('payment-form');
form.addEventListener('submit', function(event) {
  event.preventDefault();

  stripe.createToken(card).then(function(result) {
    if (result.error) {
      // Inform the customer that there was an error.
      var errorElement = document.getElementById('card-errors');
      errorElement.textContent = result.error.message;
    } else {
      // Send the token to your server.
      stripeTokenHandler(result.token);
    }
  });
});

function stripeTokenHandler(token) {
  // Insert the token ID into the form so it gets submitted to the server
  var form = document.getElementById('payment-form');
  var hiddenInput = document.createElement('input');
  hiddenInput.setAttribute('type', 'hidden');
  hiddenInput.setAttribute('name', 'stripeToken');
  hiddenInput.setAttribute('value', token.id);
  form.appendChild(hiddenInput);

  // Submit the form
  form.submit();
}