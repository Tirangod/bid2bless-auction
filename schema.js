// HTTP

// get www.bid2bless.com/ - home page
// get www.bid2bless.com/about-us - about us page
// get www.bid2bless.com/auction?id={auction_id} - certain auction page
// get www.bid2bless.com/search?request={search_request};categories={categories_names};state={auction_state};bet={last_bet_value}
// get www.bid2bless.com/create-auction //!!! user must be auth
// get www.bid2bless.com/my-auctions //!!! user must be auth

// post www.bid2bless.com/

// balance = 1000 (0 - in use) 
// ставимо 100 грн
// balance = 900 (100 - in use)
// ставку перебили
// balance = 1000 (0 - in use)

// post www.bid2bless.com/edit-auction
const edit_auction = {
    name: String,
    description: String,
    photo_bytes: Int8Array,
    start_at: Date,
    end_at: Date,
    initial_bet: Integer,
    auction_id: Integer,
    user_id: Integer
}

// post www.bid2bless.com/create-auction
const create_auction = {
    name: String,
    description: String,
    photo_bytes: Int8Array,
    start_at: Date,
    end_at: Date,
    initial_bet: Integer,
    user_id: Integer // author
}

// post www.bid2bless.com/signup - signup
const signup = {
    username: String,
    email: String,
    login_hash: Integer
}

// post www.bid2bless.com/login - login
const login = {
    login_hash: Integer
}

// WebSocket

// ws://bid2bless.com/ws/