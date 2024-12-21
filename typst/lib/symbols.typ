#let tla_symbols = (
    "\\\/": $and$,
    "(\\\/)": $or$,  
    "==": $eq.delta$,
    "CONSTANT": smallcaps("CONSTANT"),
    "ASSUME": smallcaps("ASSUME"), 
    "EXTENDS": smallcaps("EXTENDS"),
)  

#show regex("\\\/"): { 
    $or$
}
#show regex("/\\\\"): {
    $and$
}
#show regex("=="): {
    $eq.delta$
}

#show regex("CONSTANT"): {
    smallcaps("CONSTANT")
} 

#show regex("ASSUME"): { 
    smallcaps("ASSUME")
} 

#show regex("EXTENDS"): { 
    smallcaps("EXTENDS")
} 


